<template>
  <div style="padding:24px;max-width:1100px;margin:0 auto">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:20px">
      <h1 style="font-size:22px;font-weight:500">Assignments</h1>
      <button @click="showForm=true" style="background:#185fa5;color:#fff;border:none;padding:8px 16px;border-radius:6px;cursor:pointer">+ Assign Asset</button>
    </div>

    <!-- Create form -->
    <div v-if="showForm" style="position:fixed;inset:0;background:rgba(0,0,0,0.4);display:flex;align-items:center;justify-content:center;z-index:100">
      <div style="background:#fff;border-radius:12px;padding:24px;width:420px">
        <h2 style="font-size:16px;font-weight:500;margin-bottom:16px">Assign Asset to Employee</h2>
        <label style="font-size:12px;color:#888">Asset</label>
        <select v-model="form.asset_id" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
          <option value="">Select asset</option>
          <option v-for="a in assets" :key="a.id" :value="a.id">{{ a.asset_code }} — {{ a.name }}</option>
        </select>
        <label style="font-size:12px;color:#888">Employee ID (format: EMP-XXXX)</label>
        <input v-model="form.employee_id" placeholder="EMP-1001" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:10px;margin-top:4px">
        <label style="font-size:12px;color:#888">Notes (optional)</label>
        <input v-model="form.notes" placeholder="Assigned for project use" style="width:100%;padding:8px;border:0.5px solid #ddd;border-radius:6px;margin-bottom:16px;margin-top:4px">
        <div style="display:flex;gap:8px;justify-content:flex-end">
          <button @click="showForm=false" style="padding:8px 16px;border:0.5px solid #ddd;border-radius:6px;cursor:pointer;background:none">Cancel</button>
          <button @click="submit" style="padding:8px 16px;background:#185fa5;color:#fff;border:none;border-radius:6px;cursor:pointer">Assign</button>
        </div>
      </div>
    </div>

    <!-- Table -->
    <div v-if="loading" style="color:#888;font-size:14px">Loading...</div>
    <table v-else style="width:100%;border-collapse:collapse;font-size:13px">
      <thead>
        <tr style="border-bottom:0.5px solid #ddd;text-align:left">
          <th style="padding:8px">Asset ID</th>
          <th>Employee ID</th>
          <th>Assigned At</th>
          <th>Notes</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="a in assignments" :key="a.id" style="border-bottom:0.5px solid #f0f0f0">
          <td style="padding:9px 8px;font-family:monospace;color:#185fa5">{{ a.asset_id }}</td>
          <td>{{ a.employee_id }}</td>
          <td>{{ formatDate(a.assigned_at) }}</td>
          <td style="color:#888">{{ a.notes || '—' }}</td>
          <td>
            <button @click="returnAsset(a.id)" style="padding:4px 10px;background:#fcebeb;color:#a32d2d;border:none;border-radius:4px;cursor:pointer;font-size:12px">Return</button>
          </td>
        </tr>
        <tr v-if="assignments.length===0">
          <td colspan="5" style="padding:20px;text-align:center;color:#888">No active assignments</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api/axios'

const assignments = ref([])
const assets = ref([])
const loading = ref(false)
const showForm = ref(false)
const form = ref({ asset_id: '', employee_id: '', notes: '' })

async function load() {
  loading.value = true
  try {
    const [a, b] = await Promise.all([
      api.get('/assignments'),
      api.get('/assets?status=ACTIVE')
    ])
    assignments.value = a.data.data || []
    assets.value = b.data.data || []
  } finally { loading.value = false }
}

async function submit() {
  if (!form.value.asset_id || !form.value.employee_id) {
    alert('Please fill in Asset and Employee ID'); return
  }
  await api.post('/assignments', { ...form.value, assigned_by: 'admin' })
  form.value = { asset_id: '', employee_id: '', notes: '' }
  showForm.value = false
  load()
}

async function returnAsset(id) {
  if (confirm('Mark this asset as returned?')) {
    await api.put('/assignments/' + id + '/return', {})
    load()
  }
}

const formatDate = d => d ? new Date(d).toLocaleDateString() : '—'
onMounted(load)
</script>