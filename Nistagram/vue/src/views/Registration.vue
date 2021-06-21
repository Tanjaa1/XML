
<template>

  <div class="register container divs">
    <h3 class="word">New account</h3>
      <div class="row">
        <div class="col-md-6">
          <div class="form-group">
            <input id="name" type="text" class="form-control" placeholder="First Name *" value="" /><br>
            <input id="surname" type="text" class="form-control" placeholder="Last Name *" value="" />  
          </div><br>
          <div class="form-group">
            <input id="username" type="text" class="form-control" placeholder="Username *" value="" /> <br>
            <input id="email" type="email" class="form-control" placeholder="Your Email *" value="" />        
          </div><br>
          <div class="form-group">
            <input id="password1" type="password" class="form-control" placeholder="Password *" value="" /><br>
            <input id="password2" type="password" class="form-control"  placeholder="Confirm Password *" value="" /> 
          </div><br>
          <div class="form-group">         
            <input id="date" type="text" class="form-control" placeholder="Date of birth *" value="" /><br>
            <input id="phone" type="text" minlength="10" maxlength="10" name="txtEmpPhone" class="form-control" placeholder="Your Phone *" value="" />
          </div><br>
              <label class="radio inline"> 
                <input v-on:click="Gander('m')" id="male" type="radio" name="gender" value="male" checked>
                <span> Male </span> 
              </label>
              <label class="radio inline"> 
                <input v-on:click="Gander('f')" id="female" type="radio" name="gender" value="female">
                <span>Female </span> 
              </label><br>    
              <a style="color:red" v-for="e in error" :key="e">
									{{e}}<br>
							</a><br>
       <input type="submit" class="btnRegister" style="width:150px"  value="Registration" v-on:click="Registration()"/><br><br> 
          <label><small>Alredy registred?<a href="#/Login">Log in</a></small></label>    
        </div>
      </div>
  </div>
</template>

<script>
//import axios from "axios";


export default {
  name: 'Home',

  components: {
  },
  data: function () {
		return {
			error:[],
			log:false,
      account1:{
        name: null,
        surname: null,
        dateOfBirth: null,
        email: null,
        username: null,
        password: null,
        gender: null,
        phoneNumber: null
      },
      registeredUser: {
        account:{
        name: null,
        surname: null,
        dateOfBirth: null,
        email: null,
        username: null,
        password: null,
        gender: null,
        phoneNumber: null
      },
        description: null,
        website: null,
        isVerified: null,
        isPrivate: null,
        acceptingMessage: null,
        acceptingTag: null,
        userType: null
      }
		}
	},
  methods: {
		Registration(){
			this.Reset()
      this.registeredUser.account.name=document.getElementById("name").value
      this.registeredUser.account.surname=document.getElementById("surname").value
      this.registeredUser.account.username=document.getElementById("username").value
      this.registeredUser.account.password=document.getElementById("password1").value
      this.registeredUser.account.phoneNumber=document.getElementById("phone").value
      this.registeredUser.account.email=document.getElementById("email").value
      this.registeredUser.account.dateOfBirth=document.getElementById("date").value
      this.registeredUser.isVerified = false
      this.registeredUser.isPrivate = false
      this.registeredUser.acceptingMessage = false
      this.registeredUser.acceptingTag = true
      this.registeredUser.description = "xsxsxsx"
      this.registeredUser.website = ""
      if(document.getElementById("male").checked == true){
         this.registeredUser.account.gender = "MALE"    
      }else{ 
          this.registeredUser.account.gender = "FEMALE"
      }
      if(this.Validation()){
        

        fetch("http://localhost:8080/api/user/userRegistration/", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(this.registeredUser),
        mode:"no-cors"
      })
            .then(//(response) => {
              //alert(response.value)
            //if(response.status == 201){
              alert("User created")
            // }else{
            //   alert("Username not unique")
            // }
          )//})
          
          //  axios({
          //       method: "post",
          //       url: 'http://localhost:8080/api/user/userRegistration/',
          //       headers: {"Content-Type": "application/json", "crossDomain": true,"Access-Control-Allow-Origin": "*", "mode":"no-cors", credentials:"include"},
          //        data: JSON.stringify(this.registeredUser)
          //    })//.then(response => {
            //   alert(response)
            //   // if (response.status==200){
            //   //     alert('Success');
            //   //}
            // })

            // axios
            //     .post('http://localhost:8080/api/user/userRegistration/', this.registeredUser, {
            //         headers: {"Content-Type": "application/json"},
            //         mode:"no-cors"
            //     })
            //     .then(response => {
            //          alert(response.status)
            //     })
            //     .catch(error => {
            //       alert(error)
            //     })
            
      //     axios({
      //       method: "get",
      //       url:  'http://localhost:8080/api/user/getMyPersonalData/11'// + this.username,
      //   }).then(response => {
      //         if(response.status==200){
      //           this.userr = response.data;
			// 	alert(this.userr.password)
			// 	if(response.data.gender=='FEMALE')
			// 	document.getElementById("female").checked=true
			// else
			// 	document.getElementById("male").checked=true
      //         }
      //       })
      }
		},
		LogIn(){
			this.Reset()
			if(document.getElementById("logusername").value!="" && document.getElementById("logpassword").value!=""){
				//poziv
			}else{
				this.error.push('Username/password is incorrect!');
			}
		},
		Validation(){
			var r=true

			if(document.getElementById("name").value==""){
				document.getElementById("name").style.borderColor="Red"
				r=false
			}else{
				if(document.getElementById("name").value[0].match('[A-Z]'))
        this.error.push('The name may contain only letters');
      }

			if(document.getElementById("surname").value==""){
				document.getElementById("surname").style.borderColor="Red"
				r=false
			}else{
				if(document.getElementById("surname").value[0].match('[A-Z]'))
        this.error.push('The surname may contain only letters');
      }
			if(document.getElementById("password1").value==""){
				document.getElementById("password1").style.borderColor="Red"
				r=false
			}else
				this.Password(document.getElementById("password1").value,document.getElementById("password2").value)
			
			if(document.getElementById("password2").value==""){
				document.getElementById("password2").style.borderColor="Red"
				r=false
			}

			if(document.getElementById("email").value==""){
				document.getElementById("email").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("email").value.match('@gmail.com' || '@uns.ac.rs' || '@hotmail.com' || '@yahoo.com' )){
          
        this.error.push('Email form not valid!');
        }
      }

			if(document.getElementById("phone").value==""){
				document.getElementById("phone").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("phone").value.match('[0-1]'))
          this.error.push('The phone number may contain only numbers!');
      }
			if(document.getElementById("date").value==""){
				document.getElementById("date").style.borderColor="Red"
				r=false
			}else{
        if(!document.getElementById("date").value[2].match('/') || !document.getElementById("date").value[5].match('/'))
          this.error.push('Pleace put valid date form!');
      }

			if(document.getElementById("username").value==""){
				document.getElementById("username").style.borderColor="Red"
				r=false
			}else{
				if(document.getElementById("username").value[0].match('[A-Z]'))
        this.error.push('The username may contain only letters');
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
			this.error=[]
			document.getElementById("name").style.borderColor="Black"
			document.getElementById("surname").style.borderColor="Black"
			document.getElementById("password1").style.borderColor="Black"
			document.getElementById("password2").style.borderColor="Black"
			document.getElementById("email").style.borderColor="Black"
			document.getElementById("phone").style.borderColor="Black"
			document.getElementById("date").style.borderColor="Black"
			document.getElementById("username").style.borderColor="Black"
		}
	}
}
</script>

<style scoped>
.register{
    margin-top: 3%;
    padding: 3%;
    text-align: center;
}
.divs{
    border-radius: 1.5rem;
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
}
.register-left{
    text-align: center;
    color: #fff;
    margin-top: 4%;
}
.register-left input{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    width: 60%;
    background: #f8f9fa;
    font-weight: bold;
    color: #383d41;
    margin-top: 10%;
    margin-bottom: 3%;
    cursor: pointer;
}
.register-right{
    background: #f8f9fa;
    border-top-left-radius: 10% 50%;
    border-bottom-left-radius: 10% 50%;
}
.register-left img{
    margin-top: 15%;
    margin-bottom: 5%;
    width: 25%;
    -webkit-animation: mover 2s infinite  alternate;
    animation: mover 1s infinite  alternate;
}
@-webkit-keyframes mover {
    0% { transform: translateY(0); }
    100% { transform: translateY(-20px); }
}
@keyframes mover {
    0% { transform: translateY(0); }
    100% { transform: translateY(-20px); }
}
.register-left p{
    font-weight: lighter;
    padding: 12%;
    margin-top: -9%;
}
.register .register-form{
    padding: 5%;
    margin-top: 5%;
}
.register .edit-form{
    padding: 10%;
    margin-top: 10%;
}
.register-form2{
    padding: 10%;
    margin-top: 10%;
}
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #55556a;
    color: #fff;
    width: 20%;
    cursor: pointer;
}
.register .nav-tabs{
    margin-top: 3%;
    border: none;
    background: #55556a;
    border-radius: 1.5rem;
    width: 28%;
    float: right;
}
.register .nav-tabs .nav-link{
    padding: 2%;
    height: 34px;
    font-weight: 600;
    color: #fff;
    border-top-right-radius: 1.5rem;
    border-bottom-right-radius: 1.5rem;
}
.register .nav-tabs .nav-link:hover{
    border: none;
}
.register .nav-tabs .nav-link.active{
    width: 100px;
    color: #7559dc;
    border: 2px solid #9679e0;
    border-top-left-radius: 1.5rem;
    border-bottom-left-radius: 1.5rem;
}
.register-heading{
    text-align: center;
    margin-top: 8%;
    margin-bottom: -15%;
    color: #495057;
}

.word{
    color: #55556a;;
    font-weight: bold;
}

.row{
  margin-left: 40%;
}

.radio{
  margin-right: 10%;
}
</style>
