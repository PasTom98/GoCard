<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

interface Card {
  id: number
  front: string
  back: string
  lastReviewed: string | null
  nextReview: string | null
}

interface Deck {
  id: number
  title: string
  description: string
  category: string
  cardCount: number
  createdAt: string
  updatedAt: string
  createdBy: string
  views: number
  likes: number
}

const route = useRoute()
const deck = ref<Deck | null>(null)
const cards = ref<Card[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

const fetchDeckData = async () => {
  try {
    loading.value = true
    error.value = null
    
    const [deckResponse, cardsResponse] = await Promise.all([
      axios.get(`/api/decks/${route.params.id}`),
      axios.get(`/api/decks/${route.params.id}/cards`)
    ])
    
    deck.value = deckResponse.data
    cards.value = cardsResponse.data
  } catch (err) {
    error.value = 'Failed to load deck data'
    console.error('Error fetching deck data:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchDeckData()
})
</script>

<template>
  <div class="space-y-6">
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-vt-c-indigo"></div>
    </div>

    <div v-else-if="error" class="text-red-500 text-center py-12">
      {{ error }}
    </div>

    <template v-else-if="deck">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-vt-c-text-dark-1 mb-2">{{ deck.title }}</h1>
          <p class="text-vt-c-text-dark-2">{{ deck.description }}</p>
        </div>
        <div class="flex space-x-4">
          <button class="px-4 py-2 bg-vt-c-indigo text-white rounded-lg hover:bg-vt-c-indigo-dark transition-colors">
            Study Deck
          </button>
          <button class="px-4 py-2 border border-vt-c-indigo text-vt-c-indigo rounded-lg hover:bg-vt-c-indigo hover:text-white transition-colors">
            Edit Deck
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="card in cards" 
          :key="card.id"
          class="card hover:shadow-lg transition-shadow"
        >
          <div class="space-y-4">
            <div>
              <h3 class="text-lg font-semibold text-vt-c-text-dark-1 mb-2">Front</h3>
              <p class="text-vt-c-text-dark-2">{{ card.front }}</p>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-vt-c-text-dark-1 mb-2">Back</h3>
              <p class="text-vt-c-text-dark-2">{{ card.back }}</p>
            </div>
            <div class="flex justify-between items-center text-sm text-vt-c-text-dark-2">
              <span>Last reviewed: {{ card.lastReviewed ? new Date(card.lastReviewed).toLocaleDateString() : 'Never' }}</span>
              <span>Next review: {{ card.nextReview ? new Date(card.nextReview).toLocaleDateString() : 'Not scheduled' }}</span>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template> 