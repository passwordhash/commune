<script setup>
import {computed, ref} from "vue";
import InputForm from "@/components/InputForm.vue";
import router from "@/router/index.js";
import {useUserStore} from "@/stores/user.js";

const store = useUserStore()

const isAuth = ref(true);
const email = ref('')
const nickname = ref('')
const passcode = ref('')

const onsubmitLogin = (e) => {
    e.preventDefault()
    store.authorize(email.value, passcode.value)
        .then(res => {
            store.user.token = res.data.token
            if (res.status === 200) {
                router.push('/chat')
            }
        }).catch(err => alert(err));

    email.value = ""
    passcode.value = ""
}

const onsubmitRegister = (e) => {
    e.preventDefault()
    store.register(email.value, nickname.value)
        .then(res => {
            if (res.status === 200) {
                // TODO: do modal instead of alert
                alert("Вы успешно зарегистрировались!\nТеперь вы можете войти в аккунт по почте и парольному коду, высланному Вам по указанному почтовому адрессу")
                isAuth.value = true
                nickname.value = ""
            }
        })
        .catch(err => alert(err));
    // nickname.value = ""
}

const onlink = () => {
    isAuth.value = false
}

const loginClickable = computed(() => {
    return email.value.length > 0 && passcode.value.length > 0
})
const signUpClickable = computed(() => {
    return email.value.length > 0
})

</script>

<template>
    <body  class="d-flex flex-column min-vh-100">
    <div class="container container-auth text-center">
        <div class="title">
            <i class="fas fa-comments mb-4" style="font-size: 48px; color: var(--primary-color);"></i>
            <h1 class="mb-4">Commune - общайтесь без границ</h1>
        </div>
        <div class="form-wrapper row justify-content-center">
            <!-- Авторизация -->
            <Transition>
                <form class="form absolute" v-if="isAuth" id="authForm">
                    <div class="form-group">
                        <InputForm v-model="email"
                                   :placeholder="'Ваш email'"/>
                    </div>
                    <div class="form-group">
                        <InputForm v-model="passcode"
                                   :placeholder="'Парольный код'"/>
                    </div>
                    <div class="form-signup">
                        <p
                            @click="onlink"
                            class="form-link unselectable"
                        >Что такое парольный код? Регистрация</p>
                    </div>
                    <button
                        v-if="loginClickable"
                        type="submit"
                        class="btn btn-primary"
                        @click="onsubmitLogin">Войти</button>
                    <button v-else class="btn btn-primary" disabled>Войти</button>
                </form>
            </Transition>
            <!-- Регистрация -->
            <Transition>
                <form class="form absolute" v-if="!isAuth" id="signUpForm">
                    <div class="form-group">
                        <InputForm v-model="email"
                                   :placeholder="'Ваш email'"/>
                    </div>
                    <div class="form-group">
                        <InputForm v-model="nickname"
                                   :placeholder="'Ваш ник'"/>
                    </div>
                    <button
                        v-if="signUpClickable"
                        type="submit"
                        class="btn btn-primary"
                        @click="onsubmitRegister">Продолжить</button>
                    <button v-else type="submit" class="btn btn-primary" >Продолжить</button>
                </form>
            </Transition>
        </div>
    </div>
    </body>
</template>

<style>
.container-auth {
    margin: auto auto;
    padding-bottom: 200px;
}

.form-wrapper {
    position: relative;
}

.form {
    transition: all 1s;
    min-width: 50%;
    margin-left: calc(50% - 50%/2);

    position: absolute;
    top: 0;
    left: 0;
}

.form-signup {
    display: flex;
    justify-content: end;
    margin-bottom: 10px;
    margin-right: 10px;
}

.form-link {
    color: rgba(76, 175, 80, 0.55);
    font-size: 12px;
    margin-top: -10px;
    margin-bottom: 12px;
    cursor: pointer;
}

.btn-primary {
    background-color: #4CAF50;
    border-color: #4CAF50;
}

.btn-primary:disabled {
    background-color: #8BC34A;
    border-color: #8BC34A;
}

.v-enter-from {
    opacity: 0;
    translate: 200% 0;
}
.v-enter-to {
    opacity: 1;
    translate: 0 0;
}

.v-leave-from {
    opacity: 1;
    translate: 0 0;
    visibility: hidden;
}
.v-leave-to {
    opacity: 0;
    visibility: hidden;
    translate: -100% 0;
}


</style>