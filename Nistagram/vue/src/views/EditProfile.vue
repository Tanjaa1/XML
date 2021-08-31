
<template>
  <div class="divs"><small class="register-heading">* Provide your personal information,even if the account is used for a business, a pet or something else. This won't be part of public profile. *</small>			
				<br><br>
        <h3 class="register-heading word">Personal information</h3>
        
							<div class="form-group">
								First name*:
								<input id="name" v-model="userr.name"   type="text" class="form-control" disabled />
								Last name*:
								<input id="surname"  v-model="userr.surname" type="text" class="form-control" disabled />
              </div><br>
							<div class="form-group">
              Username*:
								<input id="username" v-model="userr.username" type="text" class="form-control" disabled />
							Bio:
							<input id="bio" type="text" v-model="userr.description"  class="form-control" style="width:15%" disabled />
              </div>
							<br>
							<div class="form-group">
								Email*:
								<input id="email"  v-model="userr.email" type="email" class="form-control" disabled />
								Phone number*:
								<input id="phone" v-model="userr.phoneNumber"  type="text" minlength="10" maxlength="10" name="txtEmpPhone" class="form-control" disabled />
							</div><br>
							<div class="form-group date" id="from-datepicker">
								Date of birth:
								<input id="date" v-model="userr.dateOfBirth" type="text" minlength="10" maxlength="10" class="form-control" disabled />
							Web site:
								<input id="website"  v-model="userr.website" type="text"  class="form-control" style="width:170px" disabled />
							</div><br>
              <div class="form-group">
								<div class="maxl">
									<label class="radio inline"> 
										<input v-on:click="Gander('m')" id="male" type="radio" name="gender" value="male" checked disabled >
										<span> Male </span> 
									</label>
									<label class="radio inline"> 
										<input v-on:click="Gander('f')" id="female" type="radio" name="gender" value="female" disabled >
										<span>Female </span> 
									</label>
								</div><br>
							<input type="submit" id="edit" class="btnRegister"  value="Edit" v-on:click="Edit()"/>
							<input type="submit" id="save" class="btnRegister"  value="Save" v-on:click="Save()" hidden/>
							<input type="submit" id="back" class="btnRegister"  value="Back" v-on:click="Back()" hidden/>

              <br>
              <br>
						<a style="color:red" v-for="e in error" :key="e">
							{{e}}
						</a>	
						</div><br>
					</div>
</template>

<script>
const axios=require('axios')

export default {
  name: 'EditProfile',
  components: {
  },
  data: function () {
		return {
			error:[],
			log:false,
			userr:null,
			userDto:{},
			peoples:['pera','mika','janko']
		}
	},
	beforeMount(){

		// axios({
        //     method: "get",
        //     url:  'http://localhost:8080/api/user/getMyPersonalData/20'// + this.username,
        // }).then(response => {
        //       if(response.status==200){
        //         this.userr = response.data;
		// 		alert(this.userr.password)
		// 		if(response.data.gender=='FEMALE')
		// 		document.getElementById("female").checked=true
		// 	else
		// 		document.getElementById("male").checked=true
        //       }
        //     })

			axios
                .get("http://localhost:8080/api/user/getMyPersonalData/" + localStorage.getItem('userId'),
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
					this.userr = response.data
					this.userr.dateOfBirth = this.userr.dateOfBirth.split(" ")[0]
					if(response.data.gender=='FEMALE'){
						document.getElementById("female").checked=true
					}else
						document.getElementById("male").checked=true
                }
              })

			
		
	// 	axios
    //     .get('http://localhost:8080/api/user/getMyPersonalData/5',{
	// 		headers: {
    //     "Access-Control-Allow-Origin" : "*",
    //     "Allow": "GET",
    //     "Content-type": "Application/json",
	// 	"Mode":"no-cors"
    // }})
    //     .then(response => {
    //         this.userr = response.data
	// 		alert(this.userr.name)
    //     })
    //     .catch(error => {
    //         alert(error)
    //     })
	
	// fetch('http://localhost:8080/api/user/getMyPersonalData/5', {
    //     method: "GET",
    //     headers: { "Content-Type": "application/json", "Accept": "application/json" },
    //     mode:"no-cors",
	// 	dataType: "json"
    //   })
    // .then(response => {alert(response.main)})
	// .then(data => {alert(data)})
	// //alert(this.userr)
	// this.$http.get('http://localhost:8080/api/user/getMyPersonalData/5')
	// .then(function(response){
	// 	console.log(response.data)
	// 	this.userr = response.data
	// 	alert(this.userr)
	// })
	},
	
  methods: {
		Save(){
			this.error = []
			if(this.Validation()){
			document.getElementById("name").disabled=true;
            document.getElementById("surname").disabled=true;
            document.getElementById("username").disabled=true;
            document.getElementById("bio").disabled=true;
            document.getElementById("email").disabled=true;
            document.getElementById("phone").disabled=true;
            document.getElementById("date").disabled=true;
            document.getElementById("website").disabled=true;
            document.getElementById("male").disabled=true;
            document.getElementById("female").disabled=true;
			this.userDto.name = document.getElementById("name").value,
			this.userDto.surname = document.getElementById("surname").value,
			this.userDto.dateOfBirth = document.getElementById("date").value,
			this.userDto.email = document.getElementById("email").value,
			this.userDto.username = document.getElementById("username").value,
			this.userDto.password = this.userr.password,
			
			this.userDto.phoneNumber = document.getElementById("phone").value,
			this.userDto.description = document.getElementById("bio").value,
			this.userDto.website = document.getElementById("website").value,
			
			this.userDto.isVerified=this.userr.isVerified;
			this.userDto.isPrivate=this.userr.isPrivate;
			this.userDto.acceptingMessage=this.userr.acceptingMessage;
			this.userDto.acceptingTag=this.userr.acceptingMessage;
			this.userDto.userType=this.userr.userType;
 
			if(document.getElementById("male").checked == true)
            this.userDto.gender="MALE";
			else 
			this.userDto.gender="FEMALE";


			document.getElementById("save").hidden=true;
		// 		fetch("http://localhost:8080/api/user/changeMyPersonalData/20",{
		// 	body:  JSON.stringify(this.userDto),
		// 	method: "POST",
		// 	headers: { "Content-Type": "application/json" },
		// 	mode:"no-cors"
		// })
		// 	.then(
		// 		this.userr = this.userDto
		// 	)  
			axios
                .post("http://localhost:8080/api/user/changeMyPersonalData/" + localStorage.getItem('userId'), this.userDto,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
                    alert('Successful');
					this.userr = this.userDto
                  }
              })
               .catch(error => {
                // print(error.status == 417)
                if(error == "Error: Request failed with status code 400"){
                   alert("Error")
                  }
                })
		
		// axios({
        //         method: "post",
        //         url: 'http://localhost:8080/api/user/changeMyPersonalData/5', //+ this.profileId,
        //         //headers: comm.getHeader(),
		// 		body:  JSON.stringify(this.userDto),
        //     }).then(response => {
        //       if (response.status==200){
        //           alert('Success');
        //       }
        //     })

		// 	axios({
        //     method: "put",
        //     url: "http://localhost:8080/api/user/changeMyPersonalData/10"// + this.username,
        // }).then(response => {
        //       if(response.status==200){
        //         this.userr = response.data;
		// 		alert(this.userr.name)
		// 		if(response.data.gender=='FEMALE')
		// 		document.getElementById("female").checked=true
		// 	else
		// 		document.getElementById("male").checked=true
        //       }
        //     })
			}
		},
		Edit(){
			this.Reset()
			document.getElementById("name").disabled=false;
            document.getElementById("surname").disabled=false;
            document.getElementById("username").disabled=false;
            document.getElementById("bio").disabled=false;
            document.getElementById("email").disabled=false;
            document.getElementById("phone").disabled=false;
            document.getElementById("date").disabled=false;
            document.getElementById("website").disabled=false;
            document.getElementById("male").disabled=false;
            document.getElementById("female").disabled=false;
			document.getElementById("save").hidden=false;
			document.getElementById("back").hidden=false;
			document.getElementById("edit").hidden=true;

		},Back(){
			document.getElementById("name").disabled=true;
            document.getElementById("surname").disabled=true;
            document.getElementById("username").disabled=true;
            document.getElementById("bio").disabled=true;
            document.getElementById("email").disabled=true;
            document.getElementById("phone").disabled=true;
            document.getElementById("date").disabled=true;
            document.getElementById("website").disabled=true;
            document.getElementById("male").disabled=true;
            document.getElementById("female").disabled=true;
			document.getElementById("save").hidden=true;
			document.getElementById("back").hidden=true;
			document.getElementById("edit").hidden=false;
		},
			Validation(){
			var r=true

			if(document.getElementById("name").value==""){
				document.getElementById("name").style.borderColor="Red"
				r=false
			}else{
				if(!document.getElementById("name").value[0].match('[A-Z]')){
					this.error.push('The name may contain only letters');
					r=false
				}
			}

			if(document.getElementById("surname").value==""){
				document.getElementById("surname").style.borderColor="Red"
				r=false
			}else{
				if(!document.getElementById("surname").value[0].match('[A-Z]')){
					this.error.push('The surname may contain only letters');
					r=false
				}
			}

			if(document.getElementById("email").value==""){
				document.getElementById("email").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("email").value.match('@gmail.com' || '@uns.ac.rs' || '@hotmail.com' || '@yahoo.com' )){
          r=false
        this.error.push('Email form not valid!');
        }
      }

			if(document.getElementById("phone").value==""){
				document.getElementById("phone").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("phone").value.match('[0-1]')){
          this.error.push('The phone number may contain only numbers!');
			r=false
		}
      }
			if(document.getElementById("date").value==""){
				document.getElementById("date").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("date").value[4].match('-') || !document.getElementById("date").value[7].match('-')){
          this.error.push('Pleace put valid date form!');
			r=false
		}
      }

			if(document.getElementById("username").value==""){
				document.getElementById("username").style.borderColor="Red"
				r=false
			}else{
				if(!document.getElementById("username").value[0].match('[A-Z]')){
					this.error.push('The username may contain only letters');
					r=false
				}
			}
			if(this.error==[]) return true
			else return r

		},
		Password(p1,p2){
			if(p1!=p2) this.error.push('Please confirm your password!')
		},
		Gander(g){
			if(g=='f') {
        document.getElementById("male").checked = false
      }
			else{
        document.getElementById("female").checked=false

      }

		},
		Reset(){
			this.e=[]
			document.getElementById("name").style.borderColor="Black"
			document.getElementById("surname").style.borderColor="Black"
			document.getElementById("email").style.borderColor="Black"
			document.getElementById("phone").style.borderColor="Black"
			document.getElementById("date").style.borderColor="Black"
			document.getElementById("username").style.borderColor="Black"
		}
	}
}
</script>

<style scoped>
.form-group{
  margin-left: 0%;
}
.form-control{

}
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #4608d6; 
    color: #fff;
    width: 20%;
    cursor: pointer;
}

.divs{
    min-height: 600px;
    border-radius: 1.5rem;
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
    margin-left: 3%;
    margin-right: 3%;
}
</style>
