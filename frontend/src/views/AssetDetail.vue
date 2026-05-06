<template>
  <div style="padding:24px;max-width:700px;margin:0 auto">
    <button @click="$router.back()" style="margin-bottom:16px;cursor:pointer;background:none;border:none;color:#185fa5;font-size:14px">
      ← Back
    </button>
    <div v-if="asset" style="border:0.5px solid #ddd;border-radius:12px;padding:20px">
      <div style="display:flex;justify-content:space-between;margin-bottom:16px">
        <h2 style="font-size:18px;font-weight:500">{{ asset.name }}</h2>
        <span :style="badge(asset.status)">{{ asset.status }}</span>
      </div>
      <table style="width:100%;font-size:13px;border-collapse:collapse">
        <tr v-for="row in fields" :key="row.k" style="border-bottom:0.5px solid #f0f0f0">
          <td style="padding:8px;color:#888;width:40%">{{ row.label }}</td>
          <td style="padding:8px;font-weight:500">{{ row.val }}</td>
        </tr>
      </table>
      <div style="margin-top:16px;display:flex;gap:8px">
        <button @click="editing=true" style="padding:8px 14px;border:0.5px solid #ddd;border-radius:6px;cursor:pointer">Edit</button>
        <button @click="disposeAsset" style="padding:8px 14px;background:#fcebeb;color:#a32d2d;border:none;border-radius:6px;cursor:pointer">Dispose</button>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>
<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '../api/axios'
const route = useRoute(), router = useRouter()
const asset = ref(null), editing = ref(false)
onMounted(async () => {
  const res = await api.get('/assets/'+route.params.id)
  asset.value = res.data.data
})
const fields = computed(() => !asset.value ? [] : [
  {label:'Asset Code', val:asset.value.asset_code},
  {label:'Category',   val:asset.value.category},
  {label:'Serial No.', val:asset.value.serial_number},
  {label:'Location',   val:asset.value.location},
  {label:'Purchase Value', val:'LKR '+asset.value.purchase_value?.toLocaleString()},
  {label:'Current Value',  val:'LKR '+asset.value.current_value?.toLocaleString()},
  {label:'PO ID',      val:asset.value.purchase_order_id||'—'},
])
async function disposeAsset() {
  if(confirm('Dispose this asset?')) {
    await api.delete('/assets/'+asset.value.id)
    router.push('/assets')
  }
}
const badge = s => ({
  display:'inline-block',padding:'3px 10px',borderRadius:'12px',fontSize:'12px',
  background:s==='ACTIVE'?'#eaf3de':'#fcebeb',color:s==='ACTIVE'?'#3b6d11':'#a32d2d'
})
</script>