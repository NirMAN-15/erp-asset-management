import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../api/axios'

export const useAssetStore = defineStore('assets', () => {
  const assets = ref([])
  const loading = ref(false)

  async function fetchAssets(filters = {}) {
    loading.value = true
    try {
      const res = await api.get('/assets', { params: filters })
      assets.value = res.data.data || []
    } finally { loading.value = false }
  }

  async function createAsset(data) {
    const res = await api.post('/assets', data)
    assets.value.unshift(res.data.data)
    return res.data.data
  }

  async function updateAsset(id, data) {
    const res = await api.put(`/assets/${id}`, data)
    const i = assets.value.findIndex(a => a.id === id)
    if (i > -1) assets.value[i] = res.data.data
  }

  return { assets, loading, fetchAssets, createAsset, updateAsset }
})