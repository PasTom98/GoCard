<script setup lang="ts">
const navigation = [
  { name: 'Home', to: '/' },
  { name: 'Decks', to: '/decks' },
  { name: 'Create', to: '/create' },
  { name: 'Health', to: '/health' }
]
</script>

<template>
  <div class="min-h-screen bg-[#212529] text-gray-100">
    <!-- Navigation -->
    <nav class="bg-[#1a1d20] border-b border-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <img class="h-8 w-auto" src="@/assets/logo.svg" alt="GoCard" />
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link 
                v-for="item in navigation" 
                :key="item.name"
                :to="item.to"
                class="inline-flex items-center px-1 pt-1 text-sm font-medium"
                :class="[
                  $route.path === item.to
                    ? 'text-indigo-400 border-indigo-500'
                    : 'text-gray-300 hover:text-gray-100 hover:border-gray-300'
                ]"
              >
                {{ item.name }}
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <a 
              href="/api/docs" 
              target="_blank"
              class="text-gray-300 hover:text-gray-100 px-3 py-2 text-sm font-medium"
            >
              API Docs
            </a>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
  </div>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: #1a1d20;
}

::-webkit-scrollbar-thumb {
  background: #4b5563;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #6b7280;
}
</style>
