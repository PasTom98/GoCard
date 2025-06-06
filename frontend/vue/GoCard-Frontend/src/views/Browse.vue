<template>
  <div class="space-y-6">
    <!-- Search and Filter -->
    <div class="flex flex-col md:flex-row gap-4">
      <div class="flex-1">
        <div class="relative">
          <input
            type="text"
            placeholder="Search decks..."
            class="w-full bg-secondary-800 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
          />
          <svg
            class="w-5 h-5 text-secondary-400 absolute right-3 top-2.5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            />
          </svg>
        </div>
      </div>
      <div class="flex gap-2">
        <select class="bg-secondary-800 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500">
          <option>All Categories</option>
          <option>Category 1</option>
          <option>Category 2</option>
        </select>
        <select class="bg-secondary-800 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500">
          <option>Most Recent</option>
          <option>Most Popular</option>
          <option>Most Cards</option>
        </select>
      </div>
    </div>

    <!-- Decks Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="i in 9" :key="i" class="card">
        <div class="p-4">
          <div class="flex items-center justify-between mb-4">
            <span class="text-primary-400 text-sm font-medium">Category {{ Math.ceil(i/3) }}</span>
            <button class="text-secondary-400 hover:text-white">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
              </svg>
            </button>
          </div>
          <h4 class="text-lg font-semibold text-white mb-2">Public Deck {{ i }}</h4>
          <p class="text-secondary-400 text-sm mb-4">Created by User {{ i }}</p>
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-4">
              <span class="text-secondary-300 text-sm">42 cards</span>
              <span class="text-secondary-300 text-sm">1.2k views</span>
            </div>
            <button class="text-primary-400 hover:text-primary-300 text-sm font-medium">
              View Deck
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Component logic will go here
</script> 