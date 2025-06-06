<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface HealthStatus {
  status: string
  version: string
  timestamp: string
}

const healthStatus = ref<HealthStatus | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

const checkHealth = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await axios.get('/api/health')
    healthStatus.value = response.data
  } catch (err) {
    error.value = 'Failed to fetch health status'
    console.error('Error checking health:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  checkHealth()
})
</script>

<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <div class="text-center">
      <h1 class="text-3xl font-bold text-vt-c-text-dark-1 mb-2">API Health Status</h1>
      <p class="text-vt-c-text-dark-2">Check the current status of the API</p>
    </div>

    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-vt-c-indigo"></div>
    </div>

    <div v-else-if="error" class="text-red-500 text-center py-12">
      {{ error }}
    </div>

    <div v-else-if="healthStatus" class="card space-y-4">
      <div class="flex items-center justify-between">
        <span class="text-vt-c-text-dark-1 font-medium">Status:</span>
        <span 
          :class="{
            'text-green-500': healthStatus.status === 'healthy',
            'text-red-500': healthStatus.status !== 'healthy'
          }"
        >
          {{ healthStatus.status }}
        </span>
      </div>

      <div class="flex items-center justify-between">
        <span class="text-vt-c-text-dark-1 font-medium">Version:</span>
        <span class="text-vt-c-text-dark-2">{{ healthStatus.version }}</span>
      </div>

      <div class="flex items-center justify-between">
        <span class="text-vt-c-text-dark-1 font-medium">Last Check:</span>
        <span class="text-vt-c-text-dark-2">
          {{ new Date(healthStatus.timestamp).toLocaleString() }}
        </span>
      </div>
    </div>
  </div>
</template> 