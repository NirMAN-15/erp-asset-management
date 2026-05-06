<template>
  <div style="padding:24px;max-width:1100px;margin:0 auto">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:20px">
      <h1 style="font-size:22px;font-weight:500">Maintenance</h1>
      <button @click="showForm=true" style="background:#185fa5;color:#fff;border:none;padding:8px 16px;border-radius:6px;cursor:pointer">+ Schedule</button>
    </div>

    <!-- Create form -->
    <div v-if="showForm" style="position:fixed;inset:0;background:rgba(0,0,0,0.4);display:flex;align-items:center;justify-content:center;z-index:100">
      <div style="background:#fff;border-radius:12px;padding:24px;width:420px">
        <h2 style="font-size:16px;font-weight:500;margin-bottom:16px">Schedule Maintenance</h2>
        <label style="font-size:12px;color:#888">Asset</label>
        <select v-model="form.asset_id" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
          <option value="">Select asset</option>
          <option v-for="a in assets" :key="a.id" :value="a.id">{{ a.asset_code }} — {{ a.name }}</option>
        </select>
        <label style="font-size:12px;color:#888">Description</label>
        <input v-model="form.description" placeholder="Annual service check" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
        <label style="font-size:12px;color:#888">Scheduled Date</label>
        <input v-model="form.scheduled_date" type="date" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
        <label style="font-size:12px;color:#888">Estimated Cost (LKR)</label>
        <input v-model="form.cost" type="number" placeholder="5000" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:16px;margin-top:4px">
        <div style="display:flex;gap:8px;justify-content:flex-end">
          <button @click="showForm=false" style="padding:8px 16px;border:0.5px solid #ddd;border-radius:6px;cursor:pointer;background:none">Cancel</button>
          <button @click="submit" style="padding:8px 16px;background:#185fa5;color:#fff;border:none;border-radius:6px;cursor:pointer">Schedule</button>
        </div>
      </div>
    </div>

    <!-- Table -->
    <div v-if="loading" style="color:#888;font-size:14px">Loading...</div>
    <table v-else style="width:100%;border-collapse:collapse;font-size:13px">
      <thead>
        <tr style="border-bottom:0.5px solid #ddd;text-align:left">
          <th style="padding:8px">Asset ID</th>
          <th>Description</th>
          <th>Scheduled</th>
          <th>Cost (LKR)</th>
          <th>Status</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="m in records" :key="m.id" style="border-bottom:0.5px solid #f0f0f0">
          <td style="padding:9px 8px;font-family:monospace;color:#185fa5">{{ m.asset_id }}</td>
          <td>{{ m.description }}</td>
          <td>{{ formatDate(m.scheduled_date) }}</td>
          <td>{{ m.cost ? m.cost.toLocaleString() : '—' }}</td>
          <td><span :style="badge(m.status)">{{ m.status }}</span></td>
          <td>
            <button v-if="m.status==='SCHEDULED'" @click="complete(m.id)"
              style="padding:4px 10px;background:#eaf3de;color:#3b6d11;border:none;border-radius:4px;cursor:pointer;font-size:12px">
              Complete
            </button>
          </td>
        </tr>
        <tr v-if="records.length===0">
          <td colspan="6" style="padding:20px;text-align:center;color:#888">No maintenance records</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api/axios'

const records = ref([])
const assets = ref([])
const loading = ref(false)
const showForm = ref(false)
const form = ref({ asset_id: '', description: '', scheduled_date: '', cost: '' })

async function load() {
  loading.value = true
  try {
    const [m, a] = await Promise.all([
      api.get('/maintenance'),
      api.get('/assets')
    ])
    records.value = m.data.data || []
    assets.value = a.data.data || []
  } finally { loading.value = false }
}

async function submit() {
  if (!form.value.asset_id || !form.value.scheduled_date) {
    alert('Asset and date are required'); return
  }
  await api.post('/maintenance', { ...form.value, cost: parseFloat(form.value.cost) || 0 })
  form.value = { asset_id: '', description: '', scheduled_date: '', cost: '' }
  showForm.value = false
  load()
}

async function complete(id) {
  await api.put('/maintenance/' + id + '/complete', { cost: 0 })
  load()
}

const formatDate = d => d ? new Date(d).toLocaleDateString() : '—'
const badge = s => ({
  display: 'inline-block', padding: '2px 8px', borderRadius: '12px', fontSize: '11px', fontWeight: 500,
  background: s === 'COMPLETED' ? '#eaf3de' : s === 'SCHEDULED' ? '#faeeda' : '#fcebeb',
  color: s === 'COMPLETED' ? '#3b6d11' : s === 'SCHEDULED' ? '#854f0b' : '#a32d2d'
})
onMounted(load)
</script>