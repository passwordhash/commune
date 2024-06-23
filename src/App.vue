<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Footer from "@/components/Footer.vue";
import {useUserStore} from "@/stores/user.js";

let userStore = useUserStore()

console.log(userStore.user)
</script>

<template>
  <RouterView>

  </RouterView>

  <Footer />
</template>

<style scoped>


</style>
