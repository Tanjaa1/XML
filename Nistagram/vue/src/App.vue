<template>
  <div id="app">
    <header>
    <Header/>
  </header>
  <router-view/>
  </div>
</template>

<script>
import Header from './components/Header'
export default {
  name:'app',
  components:{
    Header
  }
}
</script>
<style>
header{
  background-color: #55556a;
  padding: 15px;
}
</style>
