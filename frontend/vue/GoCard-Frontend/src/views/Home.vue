<template>
  <div class="space-y-8">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-vt-c-text-dark-1 mb-4">
        Welcome to GoCard
      </h1>
      <p class="text-vt-c-text-dark-2 text-lg">
        A modern flashcard application for efficient learning
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="card">
        <h2 class="text-2xl font-semibold text-vt-c-text-dark-1 mb-4">Create Decks</h2>
        <p class="text-vt-c-text-dark-2">
          Create and organize your flashcards into custom decks
        </p>
      </div>
      
      <div class="card">
        <h2 class="text-2xl font-semibold text-vt-c-text-dark-1 mb-4">Study Smart</h2>
        <p class="text-vt-c-text-dark-2">
          Use spaced repetition to optimize your learning
        </p>
      </div>
      
      <div class="card">
        <h2 class="text-2xl font-semibold text-vt-c-text-dark-1 mb-4">Track Progress</h2>
        <p class="text-vt-c-text-dark-2">
          Monitor your learning progress and stay motivated
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const loading = ref(false)
const error = ref<string | null>(null)
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