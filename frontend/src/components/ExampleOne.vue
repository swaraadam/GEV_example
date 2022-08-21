<template>
  <v-container color="grey lighten-5">
    <v-row no-gutters class="pb-6">
      <v-col :cols="1"> </v-col>
      <v-col :cols="10" align="right">
        <h1>Example</h1>
      </v-col>
      <v-col :cols="1"> </v-col>
    </v-row>

    <v-row>
      <v-col :cols="4"> </v-col>
      <v-col :cols="4" align="center">
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field
            v-model="name"
            :counter="10"
            :rules="nameRules"
            label="Enter your name"
            required
          ></v-text-field>

          <v-text-field
            label="Enter your e-mail address"
            v-model="email"
            :rules="emailRules"
            required
          ></v-text-field>

          <v-text-field
            label="Enter your password"
            v-model="password"
            type="password"
            :rules="passwordRules"
            required
          ></v-text-field>

          <v-btn :disabled="!valid" color="success" @click="register">
            Register
          </v-btn>
        </v-form>
      </v-col>
      <v-col :cols="4"> </v-col>
    </v-row>
    <v-snackbar v-model="snackbarActive" :timeout="2000" absolute bottom>
      {{ snackbarMessage }}
    </v-snackbar>
  </v-container>
</template>

<script>
import { RegisterRequest } from "../proto/m_user_pb";
import { UserServicePromiseClient } from "../proto/m_user_grpc_web_pb";
export default {
  data: () => ({
    valid: true,
    snackbarActive: false,
    snackbarMessage: "message!",
    name: "",
    nameRules: [
      (v) => !!v || "Name is required",
      (v) => (v && v.length <= 10) || "Name must be less than 10 characters",
    ],
    email: "",
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
    ],
    password: "",
    passwordRules: [(v) => !!v || "Password is required"],
    userDataClient: new UserServicePromiseClient(
      "http://127.0.0.1:9090",
      null,
      null
    ),
  }),
  methods: {
    async register() {
      const request = new RegisterRequest();
      request.setEmail(this.email);
      request.setPassword(this.password);
      request.setName(this.name);
      try {
        const res = await this.userDataClient.register(request, {});
        console.log(res);
        this.snackbarActive = true;
        this.snackbarMessage = res.getAccessToken();
      } catch (err) {
        console.error(err.message);
        this.snackbarActive = true;
        this.snackbarMessage = err.message;
        throw err;
      }
    },
  },
};
</script>
