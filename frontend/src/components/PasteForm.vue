<!-- frontend/src/components/PasteForm.vue -->
<template>
  <div class="paste-form">
    <div v-if="loading" class="loading">Creating paste...</div>
    <div v-if="error" class="error">{{ error }}</div>
    
    <textarea 
      v-model="content" 
      :disabled="loading"
      placeholder="Enter your text here..."
      class="paste-input"
    ></textarea>

    <button 
      @click="createPaste" 
      :disabled="!content || loading"
      class="submit-btn"
    >
      Create Paste
    </button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      content: '',
      loading: false,
      error: null,
      pasteId: null
    }
  },
  methods: {
    async createPaste() {
      this.loading = true
      this.error = null
      
      try {
        const response = await axios.post('/api/paste', {
          content: this.content,
          expire_hours: 24 // Default expiration
        })
        
        this.pasteId = response.data.id
        this.$emit('paste-created', response.data)
        this.$router.push(`/paste/${response.data.id}`)
        
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to create paste'
        console.error(error)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.paste-form {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.paste-input {
  width: 100%;
  min-height: 200px;
  margin-bottom: 20px;
  padding: 10px;
}

.submit-btn {
  padding: 10px 20px;
  background: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-btn:disabled {
  background: #cccccc;
}

.error {
  color: red;
  margin-bottom: 10px;
}

.loading {
  color: #666;
  margin-bottom: 10px;
}
</style>