<script setup>
import {computed, onBeforeMount, ref} from "vue";
import InputForm from "@/components/InputForm.vue";
import router from "@/router/index.js";
import {useUserStore} from "@/stores/user.js";

// TODO: email validation

const FORM_STATES = {
    logIn: "login",
    signUp: "signup",
    newCode: "newcod3"
}

const storeUser = useUserStore()

const formState = ref(FORM_STATES.logIn)
const email = ref('')
const nickname = ref('')
const passcode = ref('')

const formActions = {
    onsubmitLogin : (e) => {
        e.preventDefault()
        storeUser.authorize(email.value.trim(), passcode.value.trim())
            .then(res => {
                if (res.status !== 200) {
                    return
                }
                storeUser.setData(res.data.token, res.data.user)
                router.push("/chat")
            }).catch(err => alert(err.response.data.message));

        email.value = ""
        passcode.value = ""
    },
    onsubmitRegister : (e) => {
        e.preventDefault()
        storeUser.register(email.value.trim(), nickname.value.trim())
            .then(res => {
                if (res.status === 200) {
                    // TODO: do modal instead of alert
                    alert("Вы успешно зарегистрировались!\nТеперь вы можете войти в аккунт по почте и парольному коду, высланному Вам по указанному почтовому адрессу")
                    isAuth.value = true
                    nickname.value = ""
                }
            })
            .catch(err => alert(err.response.data.message));
        // nickname.value = ""
    },
    onresetCode: (e) => {
        e.preventDefault()
        storeUser.resetCode(email.value.trim())
            .then(res => {
                if (res.status === 200) {
                    // TODO: do modal instead of alert
                    alert("Вы успешно сбросили парльный код!\nТеперь вы можете войти в аккунт по почте и парольному коду, высланному Вам по указанному почтовому адрессу")
                    email.value = ""
                    formAtionsLinks.goLogIn()
                }
            })
            .catch(err => alert(err.response.data.message));
    }
}

const formAtionsLinks = {
    goSignUp: () => {
        formState.value = FORM_STATES.signUp
    },
    goNewCode: () => {
        formState.value = FORM_STATES.newCode
    },
    goLogIn: () => {
        formState.value = FORM_STATES.logIn
    }
}

const loginClickable = computed(() => {
    return email.value.trim().length > 0 && passcode.value.trim().length > 0
})
const signUpClickable = computed(() => {
    return email.value.trim().length > 0 && nickname.value.trim().length > 0
})
const newCodeClickable = computed(() => {
    return email.value.trim().length > 0
})

onBeforeMount(() => {
    if (storeUser.isAuthenticated) {
        router.push("/chat")
    }
})

</script>

<template>
    <!--    <main class="auth-main container">-->
    <main class="auth-main ">
        <div class="auth-wrapper text-center">
            <div class="title">
                <i class="fas fa-comments mb-4" style="font-size: 48px; color: var(--primary-color);"></i>
                <h1 class="mb-4">Commune - общайтесь без границ</h1>
            </div>
            <div class="form-wrapper row justify-content-center">
                <!-- Авторизация -->
                <Transition>
                    <form class="form absolute" v-if="formState === FORM_STATES.logIn" id="authForm">
                        <div class="form-group">
                            <InputForm v-model="email"
                                       :placeholder="'Ваш email'"/>
                        </div>
                        <div class="form-group">
                            <InputForm v-model="passcode"
                                       :placeholder="'Парольный код'"/>
                        </div>
                        <div class="form-subform">
                            <p
                                @click="formAtionsLinks.goNewCode"
                                class="form-link unselectable"
                            >Новый код</p>
                            <p
                                @click="formAtionsLinks.goSignUp()"
                                class="form-link unselectable"
                            >Что такое парольный код? Регистрация</p>
                        </div>
                        <button
                            v-if="loginClickable"
                            type="submit"
                            class="btn btn-primary"
                            @click="formActions.onsubmitLogin">Войти</button>
                        <button v-else class="btn btn-primary" disabled>Войти</button>
                    </form>
                </Transition>
                <!-- Регистрация -->
                <Transition>
                    <form class="form absolute" v-if="formState === FORM_STATES.signUp" id="signUpForm">
                        <div class="form-group">
                            <InputForm v-model="email"
                                       :placeholder="'Ваш email'"/>
                        </div>
                        <div class="form-group">
                            <InputForm v-model="nickname"
                                       :placeholder="'Ваш ник'"/>
                        </div>
                        <div class="form-subform">
                            <p
                                @click="formAtionsLinks.goLogIn()"
                                class="form-link unselectable"
                            >Войти в аккаунт</p>
                        </div>
                        <button
                            v-if="signUpClickable"
                            type="submit"
                            class="btn btn-primary"
                            @click="formActions.onsubmitRegister">Зарегистрироваться</button>
                        <button v-else class="btn btn-primary" disabled>Зарегистрироваться</button>
                    </form>
                </Transition>
                <!-- Новый код-->
                <Transition>
                    <form class="form absolute" v-if="formState === FORM_STATES.newCode" id="signUpForm">
                        <div class="form-group">
                            <InputForm v-model="email"
                                       :placeholder="'Ваш email'"/>
                        </div>
                        <div class="form-subform">
                            <p
                                @click="formAtionsLinks.goLogIn()"
                                class="form-link unselectable"
                            >Войти в аккаунт</p>
                        </div>
                        <button
                            type="submit"
                            class="btn btn-primary"
                            @click="formActions.onresetCode">Отправить</button>
                    </form>
                </Transition>
            </div>
        </div>
    </main>
</template>

<style>
.auth-main {
    display: flex;
    align-items: center;
    justify-content: center;
}

.auth-wrapper {
    padding-bottom: 224px;
}

.form-wrapper {
    position: relative;
}

.form {
    transition: all 1s;
    min-width: 60%;
    margin-left: calc(50% - 60%/2);

    position: absolute;
    top: 0;
    left: 0;
}

.form-subform {
    display: flex;
    justify-content: space-between;
    margin: 0 10px 10px 10px;
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