<template>
  <div style="padding:24px;max-width:1100px;margin:0 auto">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:20px">
      <h1 style="font-size:22px;font-weight:500">Asset Registry</h1>
      <button @click="showForm=true" style="background:#185fa5;color:#fff;border:none;padding:8px 16px;border-radius:6px;cursor:pointer">+ Add Asset</button>
    </div>

    <!-- Stats -->
    <div style="display:grid;grid-template-columns:repeat(4,1fr);gap:10px;margin-bottom:20px">
      <div v-for="s in stats" :key="s.label" style="background:#f5f5f3;border-radius:8px;padding:14px">
        <div style="font-size:20px;font-weight:500">{{ s.val }}</div>
        <div style="font-size:12px;color:#888">{{ s.label }}</div>
      </div>
    </div>

    <!-- Filters -->
    <div style="display:flex;gap:10px;margin-bottom:14px">
      <select v-model="filterStatus" @change="load" style="padding:6px 10px;border:0.5px solid #ddd;border-radius:6px">
        <option value="">All Status</option>
        <option>ACTIVE</option><option>MAINTENANCE</option><option>DISPOSED</option>
      </select>
      <select v-model="filterCat" @change="load" style="padding:6px 10px;border:0.5px solid #ddd;border-radius:6px">
        <option value="">All Categories</option>
        <option>IT</option><option>Furniture</option><option>Vehicle</option>
      </select>
    </div>

    <div v-if="store.loading">Loading...</div>
    <table v-else style="width:100%;border-collapse:collapse;font-size:13px">
      <thead><tr style="border-bottom:0.5px solid #ddd;text-align:left">
        <th style="padding:8px">Code</th><th>Name</th><th>Category</th>
        <th>Location</th><th>Status</th><th>Value (LKR)</th>
      </tr></thead>
      <tbody>
        <tr v-for="a in store.assets" :key="a.id"
          style="border-bottom:0.5px solid #f0f0f0;cursor:pointer"
          @click="$router.push('/assets/'+a.id)">
          <td style="padding:9px 8px;font-family:monospace;color:#185fa5">{{ a.asset_code }}</td>
          <td>{{ a.name }}</td>
          <td>{{ a.category }}</td>
          <td>{{ a.location }}</td>
          <td><span :style="badge(a.status)">{{ a.status }}</span></td>
          <td>{{ a.current_value.toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>

    <div v-if="showForm" style="position:fixed;inset:0;background:rgba(0,0,0,0.4);display:flex;align-items:center;justify-content:center;z-index:100">
      <div style="background:#fff;border-radius:12px;padding:24px;width:460px;max-height:90vh;overflow-y:auto">
        <h2 style="font-size:16px;font-weight:500;margin-bottom:16px">Add New Asset</h2>

        <label style="font-size:12px;color:#888">Asset Name *</label>
        <input v-model="newAsset.name" placeholder="Dell Laptop XPS 15" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Category *</label>
        <select v-model="newAsset.category" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
          <option value="">Select category</option>
          <option>IT</option>
          <option>Furniture</option>
          <option>Vehicle</option>
          <option>Equipment</option>
          <option>Other</option>
        </select>

        <label style="font-size:12px;color:#888">Serial Number</label>
        <input v-model="newAsset.serial_number" placeholder="SN-2024-001" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Location</label>
        <input v-model="newAsset.location" placeholder="IT Department - Floor 2" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Purchase Date *</label>
        <input v-model="newAsset.purchase_date" type="date" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Purchase Value (LKR) *</label>
        <input v-model="newAsset.purchase_value" type="number" placeholder="250000" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Asset Code *</label>
        <input v-model="newAsset.asset_code" placeholder="ASSET-0005" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">

        <label style="font-size:12px;color:#888">Purchase Order ID (from Procurement)</label>
        <input v-model="newAsset.purchase_order_id" placeholder="PO-0001" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:16px;margin-top:4px">

        <div style="display:flex;gap:8px;justify-content:flex-end">
          <button @click="showForm=false" style="padding:8px 16px;border:0.5px solid #ddd;border-radius:6px;cursor:pointer;background:none">Cancel</button>
          <button @click="submitAsset" style="padding:8px 16px;background:#185fa5;color:#fff;border:none;border-radius:6px;cursor:pointer">Create Asset</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAssetStore } from '../stores/assetStore'
const store = useAssetStore()
const filterStatus = ref(''), filterCat = ref('')
const load = () => store.fetchAssets({status:filterStatus.value,category:filterCat.value})
onMounted(load)
const showForm = ref(false)
const stats = computed(() => [
  {val: store.assets.length, label: 'Total assets'},
  {val: store.assets.filter(a=>a.status==='ACTIVE').length, label: 'Active'},
  {val: store.assets.filter(a=>a.status==='MAINTENANCE').length, label: 'In maintenance'},
  {val: 'LKR '+store.assets.reduce((s,a)=>s+a.current_value,0).toLocaleString(), label: 'Total value'},
])
const badge = s => ({
  display:'inline-block',padding:'2px 8px',borderRadius:'12px',fontSize:'11px',fontWeight:500,
  background:s==='ACTIVE'?'#eaf3de':s==='MAINTENANCE'?'#faeeda':'#fcebeb',
  color:s==='ACTIVE'?'#3b6d11':s==='MAINTENANCE'?'#854f0b':'#a32d2d'
})

const newAsset = ref({
  name: '', category: '', serial_number: '', location: '', purchase_date: '', purchase_value: '', asset_code: '', purchase_order_id: ''
})

const submitAsset = async () => {
  await store.createAsset(newAsset.value)
  showForm.value = false
  load()
}

</script>