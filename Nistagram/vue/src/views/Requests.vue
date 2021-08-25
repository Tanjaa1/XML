
<template>
  <div class="divs"><br>
                        <h3>Verification requests</h3><br>
						<div class="container">
                            <div class="row">
                                                        <table id="tableApproved" class="table table-bordered">
                                                            <thead>
                                                              <tr>
                                                                <th>Name</th>
                                                                <th>Surname</th>
                                                                <th>Category</th>
                                                                <th>Document</th>
                                                                <th>Approve</th>
                                                                <th>Disapprove</th>
                                                              </tr>
                                                            </thead>
                                                            <tbody>
                                                              <tr v-for="r in requests" :key="r">
                                                                <td>{{r.name}}</td>
                                                                <td>{{r.surname}}</td>
                                                                <td>{{r.category}}</td>
                                                                <td>{{r.photo}}</td>
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
            requests:[
                {
                    name:"Tanja",
                    surname:"Drcelic",
                    category:"BUSINESS",
                    photo:""
                },
                {
                    name:"Aleksandra",
                    surname:"Milijevic",
                    category:"BUSINESS",
                    photo:""
                }
            ],
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
   Approve(r){
       alert(r)
   },
   Disapprove(r){
       alert(r)
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
