<template>
  <div style="min-height:100vh;display:flex;align-items:center;justify-content:center">
    <div style="width:320px;border:0.5px solid #ddd;border-radius:12px;padding:24px">
      <h2 style="margin-bottom:20px;font-size:18px">ERP Asset Management</h2>
      <input v-model="username" placeholder="Username (admin)" style="width:100%;margin-bottom:10px;padding:8px;border:0.5px solid #ddd;border-radius:6px">
      <input v-model="password" type="password" placeholder="Password (admin123)" style="width:100%;margin-bottom:16px;padding:8px;border:0.5px solid #ddd;border-radius:6px">
      <button @click="doLogin" style="width:100%;padding:10px;background:#185fa5;color:white;border:none;border-radius:6px;cursor:pointer">Login</button>
      <p v-if="error" style="color:red;margin-top:10px;font-size:13px">{{ error }}</p>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/authStore'
import { useRouter } from 'vue-router'
const auth = useAuthStore(), router = useRouter()
const username = ref(''), password = ref(''), error = ref('')
async function doLogin() {
  try { await auth.login(username.value, password.value); router.push('/assets') }
  catch { error.value = 'Invalid credentials' }
}
</script>