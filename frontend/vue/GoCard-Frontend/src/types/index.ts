export interface Deck {
  id: string
  name: string
  description: string
  coverImage?: string
  cardCount: number
  createdAt: string
  updatedAt: string
  category?: string
  isPublic: boolean
  isFeatured: boolean
}

export interface Card {
  id: string
  deckId: string
  front: string
  back: string
  createdAt: string
  updatedAt: string
  lastReviewed?: string
  reviewCount: number
  difficulty: number
}

export interface Activity {
  id: string
  type: 'create' | 'update' | 'delete' | 'review'
  entityType: 'deck' | 'card'
  entityId: string
  entityName: string
  timestamp: string
  userId: string
  userName: string
} 