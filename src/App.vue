<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Footer from "@/components/Footer.vue";
import {useUserStore} from "@/stores/user.js";
import {onBeforeMount} from "vue";

const storeUser = useUserStore()

onBeforeMount(() => {
    console.log(storeUser.getAccount())
    if (storeUser.getAccount().token != null) {
        storeUser.isAuthenticated = true
    }
})

</script>

<template>
        <RouterView>

        </RouterView>

        <Footer />
</template>

<style scoped>
</style>
