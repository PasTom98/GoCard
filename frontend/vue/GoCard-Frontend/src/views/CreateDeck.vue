<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

interface Card {
  front: string
  back: string
}

const router = useRouter()
const title = ref('')
const description = ref('')
const category = ref('')
const cards = ref<Card[]>([{ front: '', back: '' }])
const loading = ref(false)
const error = ref<string | null>(null)

const addCard = () => {
  cards.value.push({ front: '', back: '' })
}

const removeCard = (index: number) => {
  if (cards.value.length > 1) {
    cards.value.splice(index, 1)
  }
}

const createDeck = async () => {
  if (!title.value || !description.value || !category.value) {
    error.value = 'Please fill in all required fields'
    return
  }

  if (cards.value.some(card => !card.front || !card.back)) {
    error.value = 'Please fill in all card fields'
    return
  }

  try {
    loading.value = true
    error.value = null

    const response = await axios.post('/api/decks', {
      title: title.value,
      description: description.value,
      category: category.value,
      cards: cards.value
    })

    router.push(`/decks/${response.data.id}`)
  } catch (err) {
    error.value = 'Failed to create deck'
    console.error('Error creating deck:', err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto space-y-8">
    <div>
      <h1 class="text-3xl font-bold text-vt-c-text-dark-1 mb-2">Create New Deck</h1>
      <p class="text-vt-c-text-dark-2">Create a new deck of flashcards to help you learn</p>
    </div>

    <form @submit.prevent="createDeck" class="space-y-6">
      <div class="space-y-4">
        <div>
          <label for="title" class="block text-sm font-medium text-vt-c-text-dark-1">Title</label>
          <input
            id="title"
            v-model="title"
            type="text"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-vt-c-indigo focus:ring-vt-c-indigo"
          />
        </div>

        <div>
          <label for="description" class="block text-sm font-medium text-vt-c-text-dark-1">Description</label>
          <textarea
            id="description"
            v-model="description"
            rows="3"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-vt-c-indigo focus:ring-vt-c-indigo"
          ></textarea>
        </div>

        <div>
          <label for="category" class="block text-sm font-medium text-vt-c-text-dark-1">Category</label>
          <input
            id="category"
            v-model="category"
            type="text"
            required
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-vt-c-indigo focus:ring-vt-c-indigo"
          />
        </div>
      </div>

      <div class="space-y-4">
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold text-vt-c-text-dark-1">Cards</h2>
          <button
            type="button"
            @click="addCard"
            class="px-4 py-2 bg-vt-c-indigo text-white rounded-lg hover:bg-vt-c-indigo-dark transition-colors"
          >
            Add Card
          </button>
        </div>

        <div v-for="(card, index) in cards" :key="index" class="card space-y-4">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-medium text-vt-c-text-dark-1">Card {{ index + 1 }}</h3>
            <button
              v-if="cards.length > 1"
              type="button"
              @click="removeCard(index)"
              class="text-red-500 hover:text-red-600"
            >
              Remove
            </button>
          </div>

          <div>
            <label :for="`front-${index}`" class="block text-sm font-medium text-vt-c-text-dark-1">Front</label>
            <textarea
              :id="`front-${index}`"
              v-model="card.front"
              rows="2"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-vt-c-indigo focus:ring-vt-c-indigo"
            ></textarea>
          </div>

          <div>
            <label :for="`back-${index}`" class="block text-sm font-medium text-vt-c-text-dark-1">Back</label>
            <textarea
              :id="`back-${index}`"
              v-model="card.back"
              rows="2"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-vt-c-indigo focus:ring-vt-c-indigo"
            ></textarea>
          </div>
        </div>
      </div>

      <div v-if="error" class="text-red-500 text-center">
        {{ error }}
      </div>

      <div class="flex justify-end space-x-4">
        <button
          type="button"
          @click="router.back()"
          class="px-4 py-2 border border-gray-300 text-vt-c-text-dark-1 rounded-lg hover:bg-gray-50 transition-colors"
        >
          Cancel
        </button>
        <button
          type="submit"
          :disabled="loading"
          class="px-4 py-2 bg-vt-c-indigo text-white rounded-lg hover:bg-vt-c-indigo-dark transition-colors disabled:opacity-50"
        >
          {{ loading ? 'Creating...' : 'Create Deck' }}
        </button>
      </div>
    </form>
  </div>
</template> 