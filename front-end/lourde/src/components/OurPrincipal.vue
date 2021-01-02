<template>
  <section>
    <div class="py-6"></div>
    <v-card
      v-show="
        principal.date == 1 &&
          !($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)
      "
      dark
      flat
    >
      <v-img class="vimage" height="600" contain :src="principal.thumbnail">
        <v-row>
          <v-col
            :cols="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 8 : 6"
            :md="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 8 : 6"
            class="ml-auto"
          >
            <v-card-text
              :style="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                  ? styled2
                  : styled
              "
              class="text-center"
            >
              Principal Address
            </v-card-text>
            <v-card-text
              :style="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                  ? styled2
                  : styled
              "
            >
              {{ principal.description }}
            </v-card-text>
            <v-row>
              <v-col md="5" cols="6" class="ml-auto">
                <v-card-text
                  :style="
                    $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                      ? styled2
                      : styled
                  "
                >
                  - {{ principal.name }}
                </v-card-text>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-img>
    </v-card>
    <v-card
      v-show="
        principal.date == 1 &&
          ($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)
      "
      dark
      flat
    >
      <v-row>
        <v-col cols="12" md="10" class="ml-auto">
          <v-card-text
            :style="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                ? styled2
                : styled
            "
            class="text-center"
          >
            Principal Address
          </v-card-text>
          <v-card-text
            :style="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                ? styled2
                : styled
            "
          >
            {{ principal.description }}
          </v-card-text>
          <v-row>
            <v-col md="5" cols="6" class="ml-auto">
              <v-card-text
                :style="
                  $vuetify.breakpoint.xs || $vuetify.breakpoint.sm
                    ? styled2
                    : styled
                "
              >
                - {{ principal.name }}
              </v-card-text>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-card>
  </section>
</template>

<script>
export default {
  name: "OurPrincipal",
  data: () => ({
    principalURL: "/api/v1/events?module=principal",
    principalData: [],
    principal: {},
    styled: "font-family:'Coronet';font-size:35px;line-height: 1.1",
    styled2: "font-family:'Coronet';font-size:19px;line-height: 1.0"
  }),
  methods: {
    async fetchData() {
      return this.axios
        .get(this.principalURL)
        .then(response => {
          return response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    this.fetchData().then(data => (this.principalData = data));
  },
  watch: {
    principalData: function(newValue) {
      if (newValue != undefined) {
        this.principal = newValue[0];
      }
    }
  }
};
</script>

<style></style>
