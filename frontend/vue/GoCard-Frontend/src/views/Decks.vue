<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-100">Your Decks</h1>
      <router-link 
        to="/create" 
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Create New Deck
      </router-link>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-900/50 border border-red-800 rounded-lg p-4 text-red-200">
      {{ error }}
    </div>

    <!-- Deck Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div 
        v-for="deck in decks" 
        :key="deck.id"
        class="bg-[#1a1d20] rounded-lg overflow-hidden border border-gray-800 hover:border-indigo-500 transition-colors duration-200"
      >
        <router-link :to="`/decks/${deck.id}`" class="block">
          <div class="aspect-w-16 aspect-h-9 bg-gray-800">
            <img 
              v-if="deck.coverImage" 
              :src="deck.coverImage" 
              :alt="deck.name"
              class="object-cover w-full h-full"
            />
            <div v-else class="flex items-center justify-center h-full">
              <span class="text-4xl text-gray-600">{{ deck.name.charAt(0) }}</span>
            </div>
          </div>
          <div class="p-4">
            <h3 class="text-lg font-semibold text-gray-100 mb-1">{{ deck.name }}</h3>
            <p class="text-sm text-gray-400 mb-2">{{ deck.description }}</p>
            <div class="flex items-center justify-between text-sm text-gray-500">
              <span>{{ deck.cardCount }} cards</span>
              <span>{{ formatDate(deck.updatedAt) }}</span>
            </div>
          </div>
        </router-link>
      </div>
    </div>

    <!-- Empty State -->
    <div v-if="!loading && !error && decks.length === 0" class="text-center py-12">
      <div class="text-gray-400 mb-4">No decks found</div>
      <router-link 
        to="/create" 
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
      >
        Create Your First Deck
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useDeckStore } from '@/stores/deck'
import type { Deck } from '@/types'

const deckStore = useDeckStore()
const loading = ref(true)
const error = ref<string | null>(null)
const decks = ref<Deck[]>([])

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

onMounted(async () => {
  try {
    loading.value = true
    decks.value = await deckStore.fetchDecks()
  } catch (e) {
    error.value = 'Failed to load decks. Please try again later.'
    console.error('Error loading decks:', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.aspect-w-16 {
  position: relative;
  padding-bottom: 56.25%;
}

.aspect-w-16 > * {
  position: absolute;
  height: 100%;
  width: 100%;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}
</style> 