<script>
export default {
  data: function() {
    return {
      error: null,
      user: {
        username: "",
        password: "",
      },
    }
  },
  methods: {
    async doLogin(e) {
      e.preventDefault()
      if (this.user.username === "") {
        this.error = "Username cannot be empty.";
      } else {
        this.error = null
        console.log(this.user)
        try {
          let response = await this.$axios.put("/login", { username: this.user.username, password: this.user.password })
          await sessionStorage.setItem("ID", response.data);
          await sessionStorage.setItem("username", this.user.username);
          this.$router.push({ path: '/success' })
        } catch (e) {
          if (e.response && e.response.status === 400) {
            this.error = "Something is wrong";
          } else if (e.response && e.response.status === 500) {
            this.error = "An internal error occurred, please try again later.";
          } else {
            this.error = e.toString();
          }
          setTimeout(() => {
            this.error = null;
          }, 5000);
        }
      }
    },

    async doSignin(e) {
      e.preventDefault()
      if (this.user.username === "") {
        this.error = "Username cannot be empty.";
      } else {
        this.error = null
        try {
          let response = await this.$axios.post("/signin", { username: this.user.username, password: this.user.password })
          await sessionStorage.setItem("ID", response.data);
          await sessionStorage.setItem("username", this.user.username);
          this.$router.push({ path: '/success' })
        } catch (e) {
          if (e.response && e.response.status === 400) {
            this.error = "Something is wrong";
          } else if (e.response && e.response.status === 500) {
            this.error = "An internal error occurred, please try again later.";
          } else {
            this.error = e.toString();
          }
          setTimeout(() => {
            this.error = null;
          }, 5000);
        }
      }
    }
  }
}
</script>

<template>
  <div class="d-flex position-relative">
    <div class="d-flex position-absolute top-0 end-0 mt-3">
      <ErrorMsg v-if="error" :msg="error"></ErrorMsg>
    </div>
  </div>

  <div>
    <div class="site-name" style="top: 10rem">
      <h1 class="h1">This is a simulation</h1>
    </div>

    <div class="d-flex justify-content-center position-absolute" style="top: 25rem; left: 0; width: 100%; height: 100%;">
      <div class="justify-content-between flex-wrap flex-md-nowrap align-items-center">
        <h2 class="h2">My pleasure to see you there</h2>
        <h2 class="h2 text-center" v-if="user.username">{{ user.username }}</h2>
      </div>
    </div>

    <div class="d-flex justify-content-center position-absolute" style="top: 30rem; left: 0; width: 100%; height: 100%; padding-top: 1rem">
      <div>
        <div class="d-flex align-items-center" style="padding-bottom: .5rem">
          <label for="username" class="form-label" style="padding-right: .5rem">Username:</label>
          <input type="text" id="username" v-model="this.user.username" class="form-control flex" style="flex-grow: 1"
                 placeholder="Insert your username" aria-label="Recipient's username" aria-describedby="basic-addon2" autocomplete="off">
        </div>
        <div class="d-flex align-items-center" style="padding-bottom: .5rem">
          <label for="password" class="form-label" style="padding-right: .5rem">Password:</label>
          <input type="password" id="password" v-model="this.user.password" class="form-control flex" style="flex-grow: 1"
                 placeholder="Insert your password" aria-label="Recipient's password" aria-describedby="basic-addon2" autocomplete="off">
        </div>
        <div class="d-flex justify-content-center">
          <div class="text-center" style="padding-top: .5rem; padding-right: 1rem">
            <button class="btn btn-outline-success" type="button" @click="doLogin">
              <div class="d-flex justify-content-between">
                <span>Log in</span>
              </div>
            </button>
          </div>
          <div class="text-center" style="padding-top: .5rem">
            <button class="btn btn-outline-dark" type="button" @click="doSignin">
              <div class="d-flex justify-content-between">
                <span>Sign in</span>
              </div>
            </button>
          </div>
        </div>

      </div>
    </div>

  </div>
</template>

<style>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
}

.site-name {
  display: flex;
  justify-content: center;
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  font-size: 3rem;
  font-weight: bold;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.site-name h1 {
  font-size: 6rem;
}

</style>
