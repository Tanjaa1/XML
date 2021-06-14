<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <ProbaRutera msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import ProbaRutera from '@/components/ProbaRutera.vue'

export default {
  name: 'Proba',
  components: {
    ProbaRutera
  }
}
</script>
