<script setup lang="ts">

</script>

<template>
  <div class="bg-vt-c-black-soft rounded-lg p-6 max-w-4xl mx-auto">
    <h2 class="text-2xl font-bold text-vt-c-text-dark-1 mb-6">System Health</h2>
    
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-vt-c-indigo"></div>
    </div>
    
    <div v-else-if="error" class="text-red-500 text-center py-12">
      {{ error }}
    </div>
    
    <div v-else class="space-y-6">
      <!-- Status and Timestamp -->
      <div class="flex justify-between items-center">
        <div class="flex items-center space-x-2">
          <div class="w-3 h-3 rounded-full" :class="status.status === 'OK' ? 'bg-green-500' : 'bg-red-500'"></div>
          <span class="text-vt-c-text-dark-1 font-medium">{{ status.status }}</span>
        </div>
        <span class="text-vt-c-text-dark-2 text-sm">{{ formatTimestamp(status.timestamp) }}</span>
      </div>

      <!-- System Metrics -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-vt-c-black-mute rounded-lg p-4">
          <h3 class="text-lg font-semibold text-vt-c-text-dark-1 mb-4">System</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">Go Version</span>
              <span class="text-vt-c-text-dark-1">{{ status.system.goVersion }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">CPU Cores</span>
              <span class="text-vt-c-text-dark-1">{{ status.system.numCPU }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">Goroutines</span>
              <span class="text-vt-c-text-dark-1">{{ status.system.numGoroutine }}</span>
            </div>
          </div>
        </div>

        <div class="bg-vt-c-black-mute rounded-lg p-4">
          <h3 class="text-lg font-semibold text-vt-c-text-dark-1 mb-4">Memory</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">Allocated</span>
              <span class="text-vt-c-text-dark-1">{{ formatBytes(status.system.memoryStats.alloc) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">Total Allocated</span>
              <span class="text-vt-c-text-dark-1">{{ formatBytes(status.system.memoryStats.totalAlloc) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">System</span>
              <span class="text-vt-c-text-dark-1">{{ formatBytes(status.system.memoryStats.sys) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-vt-c-text-dark-2">GC Cycles</span>
              <span class="text-vt-c-text-dark-1">{{ status.system.memoryStats.numGC }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Database Metrics -->
      <div v-if="status.database" class="bg-vt-c-black-mute rounded-lg p-4">
        <h3 class="text-lg font-semibold text-vt-c-text-dark-1 mb-4">Database</h3>
        <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">Max Connections</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ status.database.maxOpenConnections }}</div>
          </div>
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">Open Connections</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ status.database.openConnections }}</div>
          </div>
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">In Use</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ status.database.inUseConnections }}</div>
          </div>
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">Idle</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ status.database.idleConnections }}</div>
          </div>
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">Wait Count</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ status.database.waitCount }}</div>
          </div>
          <div class="space-y-1">
            <span class="text-vt-c-text-dark-2 text-sm">Wait Duration</span>
            <div class="text-vt-c-text-dark-1 font-medium">{{ formatDuration(status.database.waitDuration) }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { getHealth } from '../lib/api';

interface SystemMetrics {
  goVersion: string
  numCPU: number
  numGoroutine: number
  memoryStats: {
    alloc: number
    totalAlloc: number
    sys: number
    numGC: number
  }
}

interface DatabaseMetrics {
  maxOpenConnections: number
  openConnections: number
  inUseConnections: number
  idleConnections: number
  waitCount: number
  waitDuration: number
}

interface HealthResponse {
  status: string
  timestamp: string
  system: SystemMetrics
  database?: DatabaseMetrics
}

const status = ref<HealthResponse | null>(null);
const error = ref<string | null>(null);
const loading = ref(true);

const formatBytes = (bytes: number): string => {
  const units = ['B', 'KB', 'MB', 'GB']
  let value = bytes
  let unitIndex = 0

  while (value >= 1024 && unitIndex < units.length - 1) {
    value /= 1024
    unitIndex++
  }

  return `${value.toFixed(1)} ${units[unitIndex]}`
}

const formatDuration = (duration: number): string => {
  if (duration < 1000) {
    return `${duration}ns`
  } else if (duration < 1000000) {
    return `${(duration / 1000).toFixed(1)}µs`
  } else if (duration < 1000000000) {
    return `${(duration / 1000000).toFixed(1)}ms`
  } else {
    return `${(duration / 1000000000).toFixed(1)}s`
  }
}

const formatTimestamp = (timestamp: string): string => {
  return new Date(timestamp).toLocaleString()
}

const fetchHealth = async () => {
  try {
    loading.value = true
    error.value = null
    status.value = await getHealth()
  } catch (err) {
    error.value = (err as Error).message
  } finally {
    loading.value = false
  }
}

onMounted(fetchHealth)

// Refresh health data every 30 seconds
setInterval(fetchHealth, 30000)
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>