<template>
  <div class="space-y-6">
    <!-- Deck Header -->
    <div class="bg-secondary-800 rounded-lg p-6">
      <div class="flex items-start justify-between">
        <div>
          <h2 class="text-3xl font-bold text-white mb-2">Deck Name</h2>
          <p class="text-secondary-400">Created by User • Last updated 2 hours ago</p>
        </div>
        <div class="flex space-x-4">
          <button class="bg-secondary-700 hover:bg-secondary-600 text-white px-4 py-2 rounded-md text-sm font-medium">
            Edit Deck
          </button>
          <button class="bg-primary-600 hover:bg-primary-700 text-white px-4 py-2 rounded-md text-sm font-medium">
            Share
          </button>
        </div>
      </div>
    </div>

    <!-- Deck Stats -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div class="bg-secondary-800 rounded-lg p-4">
        <h3 class="text-secondary-400 text-sm font-medium">Total Cards</h3>
        <p class="text-2xl font-bold text-white mt-1">42</p>
      </div>
      <div class="bg-secondary-800 rounded-lg p-4">
        <h3 class="text-secondary-400 text-sm font-medium">Categories</h3>
        <p class="text-2xl font-bold text-white mt-1">3</p>
      </div>
      <div class="bg-secondary-800 rounded-lg p-4">
        <h3 class="text-secondary-400 text-sm font-medium">Views</h3>
        <p class="text-2xl font-bold text-white mt-1">1.2k</p>
      </div>
      <div class="bg-secondary-800 rounded-lg p-4">
        <h3 class="text-secondary-400 text-sm font-medium">Likes</h3>
        <p class="text-2xl font-bold text-white mt-1">89</p>
      </div>
    </div>

    <!-- Cards Grid -->
    <div>
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-2xl font-bold text-white">Cards</h3>
        <div class="flex space-x-4">
          <button class="text-secondary-400 hover:text-white px-3 py-2 rounded-md text-sm font-medium">
            All
          </button>
          <button class="text-secondary-400 hover:text-white px-3 py-2 rounded-md text-sm font-medium">
            Category 1
          </button>
          <button class="text-secondary-400 hover:text-white px-3 py-2 rounded-md text-sm font-medium">
            Category 2
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="i in 6" :key="i" class="bg-secondary-800 rounded-lg overflow-hidden shadow-card hover:shadow-card-hover transition-shadow duration-200">
          <div class="p-4">
            <div class="flex items-center justify-between mb-4">
              <span class="text-primary-400 text-sm font-medium">Category {{ Math.ceil(i/2) }}</span>
              <button class="text-secondary-400 hover:text-white">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                </svg>
              </button>
            </div>
            <h4 class="text-lg font-semibold text-white mb-2">Card Title {{ i }}</h4>
            <p class="text-secondary-400 text-sm mb-4">This is a sample card description that shows the content of the card.</p>
            <div class="flex items-center justify-between">
              <span class="text-secondary-300 text-sm">Last reviewed: 2h ago</span>
              <button class="text-primary-400 hover:text-primary-300 text-sm font-medium">
                Edit
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Component logic will go here
</script> 