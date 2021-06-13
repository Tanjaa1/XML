<template>
  <div class="divs"><small class="register-heading">* Provide your personal information,even if the account is used for a business, a pet or something else. This won't be part of public profile. *</small>			
				<br><br>
        <h3 class="register-heading word">Personal information</h3>
        
							<div class="form-group">
								First name*:
								<input id="name" type="text" class="form-control"/>
								Last name*:
								<input id="surname" type="text" class="form-control" />
              </div><br>
							<div class="form-group">
              Username*:
								<input id="username" type="text" class="form-control"/>
							Bio:
							<input id="bio" type="text" class="form-control" style="width:15%"/>
              </div>
							<br>
							<div class="form-group">
								Email*:
								<input id="email" type="email" class="form-control"/>
								Phone number*:
								<input id="phone" type="text" minlength="10" maxlength="10" name="txtEmpPhone" class="form-control"/>
							</div><br>
							<div class="form-group date" id="from-datepicker">
								Date of birth:
								<input id="date" type="text" minlength="10" maxlength="10" class="form-control"/>
							Web site:
								<input id="website" type="text" class="form-control" style="width:170px"/>
							</div><br>
              <div class="form-group">
								<div class="maxl">
									<label class="radio inline"> 
										<input v-on:click="Gander('m')" id="male" type="radio" name="gender" value="male" checked>
										<span> Male </span> 
									</label>
									<label class="radio inline"> 
										<input v-on:click="Gander('f')" id="female" type="radio" name="gender" value="female">
										<span>Female </span> 
									</label>
								</div><br>
							<input type="submit" class="btnRegister"  value="Edit" v-on:click="Edit()"/>
              <br>
              <br>
						<a style="color:red" v-for="e in error" :key="e">
							{{e}}
						</a>	
						</div><br>
					</div>
</template>

<script>

export default {
  name: 'EditProfile',
  components: {
  },
  data: function () {
		return {
			error:[],
			log:false
		}
	},
  methods: {
				Edit(){
			this.Reset()
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
