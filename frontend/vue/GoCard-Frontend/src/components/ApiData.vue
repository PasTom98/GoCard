<script setup lang="ts">

</script>

<template>
  <div>
    <h2>Data from Go API</h2>
    <div v-if="loading">Loading...</div>
    <pre v-else>{{ data }}</pre>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';

interface ApiResponse {
  message: string;
  // Add more fields as needed based on your Go API's JSON response
}

const data = ref<ApiResponse | null>(null);
const loading = ref(true);

onMounted(async () => {
  try {
    const response = await fetch('/api/your-endpoint');
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    data.value = await response.json() as ApiResponse;
  } catch (err) {
    console.error('Failed to fetch:', err);
    data.value = { message: 'Failed to load data' };
  } finally {
    loading.value = false;
  }
});
</script>

<style scoped>

</style>