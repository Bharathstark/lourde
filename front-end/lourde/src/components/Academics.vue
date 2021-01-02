<template>
  <section>
    <div class="py-6"></div>
    <h2 class="text-md-h3 text-sm-h3  mb-3 text-center ">Academics</h2>
    <div class="py-6"></div>
    <v-card flat>
      <v-tabs color="deep-purple accent-4" centered>
        <v-tab>12 Standard Results</v-tab>
        <v-tab>10 Standard Results</v-tab>
        <v-tab-item v-for="(students, i) in studentData" :key="i">
          <v-container>
            <v-row>
              <v-col
                v-for="student in students"
                :key="student.id"
                cols="12"
                xs="6"
                sm="4"
                md="3"
                class="d-flex flex-column"
              >
                <v-card class=" flex d-flex flex-column">
                  <div class=" mt-8 text-center">
                    <v-avatar size="130">
                      <v-img large :src="student.thumbnail"></v-img>
                    </v-avatar>
                  </div>

                  <v-list-item>
                    <v-list-item-content>
                      <v-list-item-title
                        class="text-deep-purple accent-4 ma-2 text-center"
                        >{{ student.name }}</v-list-item-title
                      >
                      <v-list-item-subtitle class=" text-center">{{
                        student.description
                      }}</v-list-item-subtitle>
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item
                    class="flex ml-14"
                    v-for="mark in student.marks"
                    :key="mark.id"
                  >
                    <v-list-item-content>{{ mark.field }}</v-list-item-content>
                    <v-list-item-content class="ml-2">
                      {{ mark.value }}
                    </v-list-item-content>
                  </v-list-item>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-tab-item>
      </v-tabs>
    </v-card>
  </section>
</template>

<script>
export default {
  name: "Academics",
  data: () => ({
    studentURL: "/api/v1/students",
    studentData: []
  }),
  created() {
    this.fetchData(this.studentURL + "?module=12").then(data =>
      this.studentData.push(data)
    );
    this.fetchData(this.studentURL + "?module=10").then(data =>
      this.studentData.push(data)
    );
  },
  methods: {
    async fetchData(url) {
      return this.axios
        .get(url)
        .then(response => {
          return response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    }
  }
};
</script>

<style></style>
