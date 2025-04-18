<template>
  <div>
    <div class="login-page-container">
      <h1>Login Page</h1>
      <form @submit.prevent="login">
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Password" required />
        <button type="submit">Login</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue' // ref'ı burada import edin
import { useRouter } from 'vue-router' // useRouter'ı burada import edin

const email = ref('')
const password = ref('')
const router = useRouter()

const login = () => {
  // Örnek olarak giriş kontrolü yapıyoruz
  if (email.value === 'user@example.com' && password.value === 'password') {
    // Başarılı girişten sonra token'ı localStorage'a kaydediyoruz
    localStorage.setItem('auth_token', 'sample_token')
    router.push('/')  // Ana sayfaya yönlendirme
  } else {
    alert('Invalid credentials')
  }
}
</script>

<style scoped>
h1 {
  font-size: 28px;
}
.login-page-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 5%;
  border: 1px solid #c8c8c8;
  border-radius: 8px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 500px;
  height: 500px;
}

form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 10px;
  width: 100%;
  padding: 0 12.5%;
}

form * {
  border: 1px solid #c8c8c8;
  border-radius: 8px;
  width: 100% !important;
  padding: 8px 15px;
}

form > button {
  background-color: #077A7D;
  color: #fff;
  transition: background-color 0.5s ease-in-out;
}

form > button:hover {
  background-color: #7AE2CF;
}
</style>
