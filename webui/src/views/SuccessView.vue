<script>
export default {
  data: function() {
    return {
      error: null,
      username: "",
      ID: 0,
    }
  },
  methods: {
    async retry() {
      await sessionStorage.clear()
      this.$router.push({path: '/'})
    },
  }
}
</script>

<template>
  <div class="d-flex position-relative">
    <div class="d-flex position-absolute top-0 end-0 mt-3">
      <ErrorMsg v-if="error" :msg="error"></ErrorMsg>
    </div>
  </div>

  <div></div>

  <div>
    <div class="d-flex justify-content-center position-absolute" style="top: 45%; left: 0; width: 100%; height: 100%;">
      <div class="justify-content-between flex-wrap flex-md-nowrap align-items-center">
        <h2 class="h1">Operation has been done successfully</h2>
      </div>
    </div>

    <div class="d-flex justify-content-center align-items-center position-absolute" style="left: 0; width: 100%; height: 100%; padding-top: 1rem;">
      <div class="text-center">
        <button type="button" class="btn" style="border: none;" @click="retry">
            <span class="site-name" style="font-size: 1.25rem; font-weight: normal; color: red">Retry</span>
        </button>
      </div>
    </div>
  </div>
</template>

<style>
</style>
