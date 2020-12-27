<template>
  <v-snackbar :timeout="timeout" v-model="snackbar" :color="color" multi-line>
    {{ message }}

    <template v-slot:action="{ attrs }">
      <v-btn text v-model="showLogin" to="/login">Login</v-btn>
      <v-btn color="red" text v-bind="attrs" @click="snackbar = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
export default {
  name: "showed",
  data: () => ({
    snackbar: false,
    showLogin: false,
    timeout: 10000,
    message: "",
    color: ""
  }),
  methods: {
    open(message, color, timeout, login) {
      this.snackbar = true;
      this.color = color;
      this.showLogin = login;
      this.message = message;
      this.timeout = timeout;
    }
  }
};
</script>
