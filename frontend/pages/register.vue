<template>
    <div>
      <div class="register-page-container">
        <h1>Register Page</h1>
        <form @submit.prevent="register">
            <input v-model="name" type="text" placeholder="Name" required />
            <input v-model="surname" type="text" placeholder="Surname" required />
            <input v-model="email" type="email" placeholder="Email" required />
            <input v-model="password" type="password" placeholder="Password" required />
            <button type="submit">Sign Up</button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  
  const email = ref('')
  const password = ref('')
  const router = useRouter()
  
  const login = () => {
    if (email.value === 'user@example.com' && password.value === 'password') {
      localStorage.setItem('auth_token', 'sample_token')
      router.push('/')
    } else {
      alert('Invalid credentials')
    }
  }
  </script>
  
  <style scoped>
  h1 {
    font-size: 28px;
  }
  .register-page-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 5%;
    border: 1px solid #c8c8c8;
    border-radius: 8px;
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