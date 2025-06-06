import { defineStore } from 'pinia'
import axios from 'axios'

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

interface DeckState {
  decks: Deck[]
  featuredDecks: Deck[]
  recentActivity: any[]
  loading: boolean
  error: string | null
}

export const useDeckStore = defineStore('deck', {
  state: (): DeckState => ({
    decks: [],
    featuredDecks: [],
    recentActivity: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchDecks() {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get('/api/decks')
        this.decks = response.data
      } catch (err) {
        this.error = 'Failed to fetch decks'
        console.error('Error fetching decks:', err)
      } finally {
        this.loading = false
      }
    },

    async fetchFeaturedDecks() {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get('/api/decks/featured')
        this.featuredDecks = response.data
      } catch (err) {
        this.error = 'Failed to fetch featured decks'
        console.error('Error fetching featured decks:', err)
      } finally {
        this.loading = false
      }
    },

    async fetchRecentActivity() {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get('/api/activity')
        this.recentActivity = response.data
      } catch (err) {
        this.error = 'Failed to fetch recent activity'
        console.error('Error fetching recent activity:', err)
      } finally {
        this.loading = false
      }
    }
  }
}) 