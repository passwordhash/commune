<script setup>
import {computed, ref} from "vue";
import InputForm from "@/components/InputForm.vue";
import router from "@/router/index.js";

const name = ref('')
const phone = ref('')
const code = ref('')
const submit = ref(false)

const onsubmit = (e) => {
  e.preventDefault()
  router.push('/chat')
}

const clickable = computed(() => {
  return true
  // return name.value.length > 0 && phone.value.length > 0 && code.value.length > 0;
})

</script>

<template>
  <body class="d-flex flex-column min-vh-100">
  <div class="container my-auto text-center">
    <i class="fas fa-comments mb-4" style="font-size: 48px; color: #4CAF50;"></i>
    <h1 class="mb-4">Commune - общайтесь без границ</h1>
    <div class="row justify-content-center">
      <div class="col-sm-10 col-md-8 col-lg-6">
        <form id="authForm">
          <div class="form-group">
            <InputForm
                :placeholder="'Номер телефона'"
                v-model="phone"/>
          </div>
          <div class="form-group">
            <InputForm v-model="code"
                       :placeholder="'Код подтверждения'"/>
          </div>
          <div class="form-group">
            <InputForm v-model="name"
                       :placeholder="'Выше имя'"/>
          </div>
          <button
              v-if="clickable"
              type="submit"
              class="btn btn-primary"
              @click="onsubmit">Продолжить</button>
          <button v-else type="submit" class="btn btn-primary" disabled>Продолжить</button>
        </form>
      </div>
    </div>
  </div>


  </body>
</template>

<style>
.btn-primary {
  background-color: #4CAF50;
  border-color: #4CAF50;
}

.btn-primary:disabled {
  background-color: #8BC34A;
  border-color: #8BC34A;
}

.form-control {
  border-radius: 20px;
  border-color: #4CAF50;
  margin-bottom: 10px;
  height: 50px;
//max-width: 400px;
}

.form-control:focus {
  border-color: #4CAF50;
  box-shadow: 0 0 0 0.2rem rgba(76, 175, 80, 0.25);
}
</style>