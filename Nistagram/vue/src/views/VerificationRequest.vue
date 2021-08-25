
<template>
  <div class="divs"><br>
							<div class="form-group">
                           <h3 >Verification request</h3><br>
								First name*:
								<input id="name"  type="text" class="form-control" placeholder="Full first name" />
								Last name*:
								<input maxlength="4" id="surname"   type="text" class="form-control" placeholder="Full last surname"/>
                            </div><br>
							<div class="form-group">
								Category*:
								<select name="category" id="category">
                                <optgroup label="Category">
                                    <option value="INFLUENCER">INFLUENCER</option>
                                    <option value="NEWSMEDIA">NEWSMEDIA</option>
                                    <option value="BUSINESS">BUSINESS</option>
                                    <option value="BRAND">BRAND</option>
                                    <option value="ORGANIZATION">ORGANIZATION</option>
                                    <option value="ARTIST">ARTIST</option>
                                    <option value="EDUCATION">EDUCATION</option>
                                </optgroup>                              
                                </select><br><br>
                                Upload document*:
                                <input type="file" @change="onFileChange($event)"><br>
                                <img :src="image" style="width:25%"/>
							</div><br>
                            <input style="margin-left:20%" type="submit" class="btnRegister"  value="Sent" v-on:click="Sent()"/>
                            <a style="color:red">
							{{error}}
						</a>	
					</div>
</template>

<script>
const axios=require('axios')

export default {
  name: 'VerificationRequest',
  components: {
  },
  data: function () {
		return {
            image:null,
            error:'',
		}
	},
	beforeMount(){
	},
	
  methods: {
      
   onFileChange( e) {

       const URL = 'http://foobar.com/upload'; 

    let data = new FormData();
    data.append('name', 'my-picture');
    data.append('file', e.target.files[0]); 

    let config = {
      header : {
        'Content-Type' : 'image/png'
      }
    }

    axios.put(
      URL, 
      data,
      config
    ).then(
      response => {
        console.log('image upload response > ', response)
      }
    )

      var files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      this.createImage(files[0]);
    },
    createImage(file) {
      var reader = new FileReader();

      reader.onload = (e) => {
        this.image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    Sent(){
        if(this.image!=null && document.getElementById("surname").value!="" && document.getElementById("name").value!=""
        && document.getElementById("categoty").value!="")
        {
            //KOD
        }else
        this.error='You must fill in all the fields!'
	}
}
}
</script>

<style scoped>
.form-group{
  margin-left: 20%;
  margin-right: 20%;
}
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #55556a;
    color: #fff;
    width: 10%;
    cursor: pointer;
}

.divs{
    min-height: 300px;
    border-radius: 1.5rem;
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
    margin-left: 3%;
    margin-right: 3%;  
    align-items: center;
    
}
</style>
