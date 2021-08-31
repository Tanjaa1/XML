
<template>
  <div class="divs"><br>
                        <h3>Verification requests</h3><br>
						<div class="container">
                            <div class="row">
                                                        <table id="tableApproved" class="table table-bordered">
                                                            <thead>
                                                              <tr>
                                                                <th>Name Surname</th>
                                                                <th>Username</th>
                                                                <th>Approve</th>
                                                                <th>Disapprove</th>
                                                              </tr>
                                                            </thead>
                                                            <tbody>
                                                              <tr v-for="r in requests" :key="r">
                                                                <td>{{r.user}}</td>
                                                                <td>{{r.username}}</td>
                                                                <td style="text-align:center"><button class="btnRegister" v-on:click="Approve(r)">A P P R O V E</button></td> 
                                                                <td style="text-align:center"><button class="btnRegister" v-on:click="Disapprove(r)">D I S A P P R O V E</button></td>  
                                                              </tr>
                                                            </tbody>
                                                        </table>
                                                        </div>
                                              </div>		
					</div>
</template>

<script>
const axios=require('axios')

export default {
  name: 'Requests',
  components: {
  },
  data: function () {
		return {
            requests:[],
            image:null,
            error:'',
            userr:null
		}
	},
	beforeMount(){

         axios
                .get("http://localhost:8080/api/user-request/getRequestsByReceiverId/" + localStorage.getItem('userId'),
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
                    this.requests = response.data
                    //alert("Usao u metodu")
                    //alert(this.requests.length)
                    if(this.requests != null){
                      //alert("aa")
                      for(let i = 0; i < this.requests.length; i++){
                        //alert( this.requests[i].fol1)
                          axios
                                  .get("http://localhost:8080/api/user/getMyPersonalData/" + this.requests[i].fol1,
                          {
                                headers: {
                                  'Authorization': 'Bearer' + " " + localStorage.getItem('token')
                                }
                          })
                                  .then(response1 => {
                                    if (response1.status==200){
                                      this.userr = response1.data
                                        if(this.userr != null){
                                          this.requests[i].user = this.userr.name + this.userr.surname
                                          this.requests[i].username = this.userr.username
                                        }
                                    }
                                })
                      }
                    }
                  
                }
              })

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
   Approve(r){
     
       let c = {
          fol1: parseInt(3),
          fol2: parseInt(2),
        }
       axios
                .put("http://localhost:8080/api/user/addFollower/" + localStorage.getItem('userId') + '/' + r.fol1,c,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==200){
                    this.Disapprove(r)
                    location.reload()
                  }
              })
               .catch(error => {
                // print(error.status == 417)
                if(error == "Error: Request failed with status code 400"){
                   alert("Error")
                  }
                })

   },
   Disapprove(r){
 
       axios
                                  .get("http://localhost:8080/api/user-request/deleteRequest/" + r.id,
                          {
                                headers: {
                                  'Authorization': 'Bearer' + " " + localStorage.getItem('token')
                                }
                          })
                                  .then(response1 => {
                                    if (response1.status==200){
                                      location.reload()
                                    }
                                })
   }
}
}
</script>

<style scoped>
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #55556a;
    color: #fff;
    cursor: pointer;
}

.divs{
    min-height: 300px;
    border-radius: 1.5rem;
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
    margin-left: 3%;
    margin-right: 3%;  
    text-align:center
}
.tableBorder {
    border: hidden;
    border-width: medium;
    border-collapse: separate;
    border-spacing: 0 15px;
    font-family: Impact;
    background-image: linear-gradient(to right, lightblue,white);
}
</style>
