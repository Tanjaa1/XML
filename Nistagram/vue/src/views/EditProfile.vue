
<template>
  <div class="divs"><small class="register-heading">* Provide your personal information,even if the account is used for a business, a pet or something else. This won't be part of public profile. *</small>			
				<br><br>
        <h3 class="register-heading word">Personal information</h3>
        
							<div class="form-group">
								First name*:
								<input id="name" v-model="userDto.name"   type="text" class="form-control" disabled />
								Last name*:
								<input id="surname"  v-model="userDto.surname" type="text" class="form-control" disabled />
              </div><br>
							<div class="form-group">
              Username*:
								<input id="username" v-model="userDto.username" type="text" class="form-control" disabled />
							Bio:
							<input id="bio" type="text" v-model="userDto.description"  class="form-control" style="width:15%" disabled />
              </div>
							<br>
							<div class="form-group">
								Email*:
								<input id="email"  v-model="userDto.email" type="email" class="form-control" disabled />
								Phone number*:
								<input id="phone" v-model="userDto.phone"  type="text" minlength="10" maxlength="10" name="txtEmpPhone" class="form-control" disabled />
							</div><br>
							<div class="form-group date" id="from-datepicker">
								Date of birth:
								<input id="date" v-model="userDto.dateOfBirth" type="text" minlength="10" maxlength="10" class="form-control" disabled />
							Web site:
								<input id="website"  v-model="userDto.website" type="text"  class="form-control" style="width:170px" disabled />
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
							<input type="submit" class="btnRegister"  value="Edit" v-on:click="Edit()"/>
							<input type="submit" id="save" class="btnRegister"  value="Save" v-on:click="Save()" hidden/>
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
			userDto:{}

		}
	},
	beforeMount(){

		axios
        .get('http://localhost:8080/api/user/getMyPersonalData/20/',{
			headers: {
        "Access-Control-Allow-Origin" : "",
        "Allow": "GET",
        "Content-type": "Application/json",
		"Mode":"no-cors"
    }})
        .then(response => {
            this.userr = response.data
			alert(this.userr.name)
        })
        .catch(error => {
            alert(error)
        })
	},
	
  methods: {
		Save(){
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
			this.userDto.password = document.getElementById("password").value,
			this.userDto.phoneNumber = document.getElementById("phone").value,
			this.userDto.description = document.getElementById("bio").value,
			this.userDto.website = document.getElementById("website").value,
			
			this.userDto.isVerified=this.userr.isVerified;
			this.userDto.isPrivate=this.userr.isPrivate;
			this.userDto.acceptingMessage=this.userr.acceptingMessage;
			this.userDto.acceptingTag=this.userr.acceptingMessage;;
			this.userDto.userType=this.userr.userType;

    
 
			if(document.getElementById("male").checked == true)
            this.userDto.gender="MALE";
			else 
			this.userDto.gender="FEMALE";


			document.getElementById("save").hidden=true;
				fetch("http://localhost:8080/api/user/changeMyPersonalData/20",{
			body:  JSON.stringify(this.userDto),
			method: "PUT",
			headers: { "Content-Type": "application/json" },
			mode:"no-cors"
		})
			.then(response => {
			this.userr = response.data
			})    
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

			if(this.Validation()){
				//poziv
			} 
		},
		Validation(){
			var r=true

			if(document.getElementById("name").value==""){
				document.getElementById("name").style.borderColor="Red"
				r=false
			}else
				this.Latters(document.getElementById("name").value)

			if(document.getElementById("surname").value==""){
				document.getElementById("surname").style.borderColor="Red"
				r=false
			}else
				this.Latters(document.getElementById("surname").value)

			if(document.getElementById("email").value==""){
				document.getElementById("email").style.borderColor="Red"
				r=false
			}

			if(document.getElementById("phone").value==""){
				document.getElementById("phone").style.borderColor="Red"
				r=false
			}
			if(document.getElementById("date").value==""){
				document.getElementById("date").style.borderColor="Red"
				r=false
			}

			if(document.getElementById("username").value==""){
				document.getElementById("username").style.borderColor="Red"
				r=false
			}else
				this.Latters(document.getElementById("username").value)

			if(this.error==[]) return true
			else return r

		},
		Latters(str){		
			let nameMatch = str.match('[A-Za-z ]*');
			if (nameMatch != str) this.error.push('The name may contain only letters');
			else if (str[0].match('[A-Z]') === null) this.error.push('The name must begin with a capital letter');
		},
		Password(p1,p2){
			if(p1!=p2) this.error.push('Please confirm your password!')
		},
		Gander(g){
			if(g=='f')
				document.getElementById("male").checked=false
			else
				document.getElementById("female").checked=false
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
