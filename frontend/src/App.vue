<template>
  <div v-if="!isLogin" style="display:flex;min-height:100vh">
    <nav style="width:200px;border-right:0.5px solid #ddd;padding:20px 0;flex-shrink:0">
      <div style="padding:0 16px 20px;font-weight:500;font-size:15px">ERP Assets</div>
      <router-link v-for="l in links" :key="l.path" :to="l.path"
        style="display:block;padding:10px 16px;font-size:13px;text-decoration:none;color:inherit"
        active-class="nav-active">{{ l.label }}
      </router-link>
      <div @click="logout" style="padding:10px 16px;font-size:13px;cursor:pointer;color:#a32d2d;margin-top:auto">Logout</div>
    </nav>
    <main style="flex:1;overflow:auto"><router-view/></main>
  </div>
  <router-view v-else/>
</template>
<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from './stores/authStore'
const route = useRoute(), router = useRouter(), auth = useAuthStore()
const isLogin = computed(() => route.path === '/login')
const links = [{path:'/assets',label:'Assets'},{path:'/assignments',label:'Assignments'},{path:'/maintenance',label:'Maintenance'}]
function logout() { auth.logout(); router.push('/login') }
</script>
<style>
.nav-active { background:#e6f1fb; color:#185fa5 !important; border-radius:6px; }
* { box-sizing: border-box; font-family: sans-serif; }
</style>