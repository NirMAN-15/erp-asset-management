<template>
  <div style="display:flex;height:100vh;align-items:center;justify-content:center;background:#f5f5f3">
    <div style="background:#fff;padding:32px;border-radius:12px;width:360px;box-shadow:0 4px 12px rgba(0,0,0,0.05)">
      <h1 style="font-size:24px;font-weight:500;margin-bottom:24px;text-align:center">Login</h1>
      <form @submit.prevent="login">
        <label style="display:block;font-size:13px;margin-bottom:6px">Username</label>
        <input v-model="username" style="width:100%;padding:10px;border:1px solid #ddd;border-radius:6px;margin-bottom:16px">
        <label style="display:block;font-size:13px;margin-bottom:6px">Password</label>
        <input v-model="password" type="password" style="width:100%;padding:10px;border:1px solid #ddd;border-radius:6px;margin-bottom:24px">
        <button type="submit" style="width:100%;background:#185fa5;color:#fff;border:none;padding:12px;border-radius:6px;cursor:pointer;font-weight:500">Log In</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/authStore'

const username = ref('')
const password = ref('')
const router = useRouter()
const auth = useAuthStore()

const login = async () => {
  try {
    await auth.login(username.value, password.value)
    router.push('/')
  } catch (err) {
    alert('Login failed')
  }
}
</script>
