<template>
  <section>
    <div class="py-6"></div>
    <h2 id="ourSchoolHeader" class="text-md-h3 text-sm-h3  mb-3 text-center ">
      Our School
    </h2>
    <v-row no-gutters cols="12" md="6">
      <v-scale-transition>
        <v-col
          v-show="!($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)"
          cols="4"
        >
          <v-img max-height="80%" src="/img/slider/mother.jpg"> </v-img>
        </v-col>
      </v-scale-transition>
      <v-col>
        <v-container v-for="(item, i) in ourSchool" :key="i">
          <v-card flat>
            <v-card-title>{{ item.title }}</v-card-title>
            <v-card-text class="text-md-body-1">{{ item.text }}</v-card-text>
          </v-card>
        </v-container>
        <v-card-actions>
          <v-btn text class=" ml-4" href="http://www.cmcsisters.org/" target="_blank" color="deep-purple accent-4">
            Learn More
          </v-btn>
        </v-card-actions>
      </v-col>
    </v-row>
  </section>
</template>

<script>
export default {
  name: "OurSchool",
  data: () => ({
    ourSchool: [
      {
        title: "Our Vision",
        text:
          "We aim at an all round formation where the character and personality development of the student, based on the teachings of Christ is focused at, so as to make them useful and efficient members of the society and the world at large "
      },
      {
        title: "Our Mission",
        text:
          "Our Mission is to educate children to become active and responsible citizens. To impart knowledge and wisdom with love and dedication. To mould Christ like leaders who will transform the society"
      },
      {
        title: "Our Establishment",
        text:
          "Our School is recognized by the board of Matriculation school under the Education Department of Tamil Nadu and trains students for Matriculation and Higher Secondary Examinations. The medium of instruction and communication is English.Currently Lourde Matha academic community consists of 2000 plus students and 100 staffs. Both the staff and students are orienting for an integrated life with a vision of self formation. Staff are dedicated for the student community, for their scholastic excellence in academics and in other activities. The commitment of the staff, co-operation of the parents, response of the students, altogether makes the institution vibrant and excellent.The management sees that the institution not only imparts knowledge but also cares for spiritual and moral well being of pupils, entrusted under their care."
      }
    ]
  })
};
</script>

<style></style>
