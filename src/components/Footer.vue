<script setup>
import {useUserStore} from "@/stores/user.js";
import {computed} from "vue";

const storeUser = useUserStore()

const onLogout = () => {
    console.log("logout")
    storeUser.logout()
}

</script>

<template>
    <footer :style="{
        // '--primary-color': store.primaryColor,
        // '--background-color': store.primaryColor
    }"  class="footer mt-auto py-3 d-flex">
        <div class="container my-auto d-flex">
            <a class="d-flex align-items-center mr-auto" href="https://github.com/passwordhash" target="_blank">
                <span class="text-muted">Сделано автором</span>
                <i class="fab text-muted fa-github pl-2"></i>
            </a>
            <button
                v-if="storeUser.isAuthenticated"
                type="button"
                @click="onLogout"
                class="btn btn-secondary mr-2"
            >Выйти</button>
        </div>
    </footer>
</template>

<style scoped>
a {
    text-decoration: none;
}
.footer {
    background-color: var(--background-color);
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    height: 70px;
}
</style>