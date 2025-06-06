import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/decks',
      name: 'decks',
      component: () => import('../views/Decks.vue')
    },
    {
      path: '/decks/:id',
      name: 'deck-detail',
      component: () => import('../views/DeckDetail.vue')
    },
    {
      path: '/create',
      name: 'create-deck',
      component: () => import('../views/CreateDeck.vue')
    },
    {
      path: '/health',
      name: 'health',
      component: () => import('../views/HealthCheck.vue')
    }
  ]
})

export default router 