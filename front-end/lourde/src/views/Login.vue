<template>
  <v-container>
    <v-layout class="pt-15"> </v-layout>
    <v-layout class="pt-10"> </v-layout>
    <v-layout wrap>
      <v-flex class="pt-15" sm8 md6 offset-md3 offset-sm2>
        <v-card elevation="4" light tag="section">
          <v-card-title>
            <v-layout align-center justify-space-between>
              <h3 class="headline">
                Admin Login
              </h3>
              <v-flex>
                <v-img
                  class="shrink"
                  position="center right"
                  src="../assets/logos.png"
                  contain
                  height="50px"
                />
              </v-flex>
            </v-layout>
          </v-card-title>
          <v-divider></v-divider>
          <v-form ref="loginForm">
            <v-card-text>
              <span class="caption grey--text text--darken-1">
                Enter the user to login to your Lourde account
              </span>
              <v-text-field
                outlined
                label="User"
                :rules="[v => !!v || 'field is required']"
                v-model="username"
              ></v-text-field>
              <v-text-field
                outlined
                :rules="[v => !!v || 'field is required']"
                label="Password"
                type="password"
                v-model="password"
              ></v-text-field>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions :class="{ 'pa-3': $vuetify.breakpoint.smAndUp }">
              <v-spacer></v-spacer>
              <v-btn
                color="info"
                :loading="loading"
                :large="$vuetify.breakpoint.smAndUp"
                @click="login"
              >
                Login
                <v-icon left>mdi-lock</v-icon>
              </v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-flex>
    </v-layout>
    <alertDialog ref="showed"></alertDialog>
  </v-container>
</template>

<script>
import alertDialog from "@/components/SnackBarDialog.vue";
export default {
  name: "Login",
  components: {
    alertDialog
  },
  data: () => ({
    step: 1,
    loading: false,
    username: "",
    password: "",
    apiURL: "/api/auth/login",
    name: {},
    data: {}
  }),
  methods: {
    login() {
      if (this.$refs.loginForm.validate()) {
        this.loading = true;
        const params = new URLSearchParams();
        params.append("username", this.username);
        params.append("password", this.password);
        this.axios
          .post(this.apiURL, params, {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          })
          .then(response => {
            if (response.data.response_status.response_code == 200) {
              this.loading = false;
              let data = response.data["data"];
              this.$cookies.set("csrf_token", data.token, 0);
              this.$cookies.set("username", data.name, 0);
              this.$router.push("/admin");
              this.$refs.showed.open(
                "Logged in Successfully",
                "blue",
                2000,
                false
              );
            } else {
              this.loading = false;
              this.$refs.showed.open(
                "Login Failed - User Name Password Error",
                "red",
                2000,
                false
              );
            }
          });
      }
    }
  }
};
</script>

<style></style>
