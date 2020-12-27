<template>
  <v-main>
    <v-toolbar color="white" flat>
      <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
      <AppBar v-model="toggle"></AppBar>

      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title class="grey--text text--darken-4">
        Contact Us
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        v-for="links in icons"
        :key="links.icon"
        :href="links.link"
        target="_blank"
        class="mx-4"
        icon
      >
        <v-icon size="24px">
          {{ links.icon }}
        </v-icon>
      </v-btn>
    </v-toolbar>
    <section>
      <v-card flat width="100%" class="lighten-1 text-center">
        <div class="py-6"></div>
        <v-img class="shrink" src="../assets/logos.png" contain height="80px" />
        <v-card-text dark class="text-md-h5 font-weight"
          >Lourde Matha Convent Matriculation Higher Secondary School
          <br />Dhali road, Bodipatti(PO), <br />Udumalpet - 642154, <br />Tamil
          Nadu,
        </v-card-text>
      </v-card>
    </section>
    <section>
      <div class="py-6"></div>

      <v-row>
        <v-col cols="12" md="8">
          <v-card class="ml-10 mb-10" flat>
            <v-card-title>
              Location
            </v-card-title>
            <iframe
              allowfullscreen=""
              frameborder="0"
              height="450"
              src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d15689.504257428734!2d77.2284804!3d10.5497163!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0x5912edd2fa044d86!2sLourde+Matha+Convent+Matric.+Hr.+Sec.+School!5e0!3m2!1sen!2sin!4v1539448496105"
              style="border:0"
              width="100%"
            ></iframe>
          </v-card>
        </v-col>
        <v-col>
          <v-card class="ma-16" flat>
            <v-card-title>
              Working Hours
            </v-card-title>
            <v-card-text class="text-subtitle-1">
              9.00 am to 3.15 pm (Monday to Friday)<br />
              9.00 am to 1.00 noon (Working Saturday)<br />
              School office will remain closed on Sundays and Public
              Holidays.<br />
            </v-card-text>
            <v-card-title>
              Contact Details
            </v-card-title>
            <v-card-text class="text-subtitle-1">
              04252 - 231288<br />
              94866 16232<br />
              lmcm.school@yahoo.com<br />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </section>
    <v-card flat>
      <v-card-text>
        Copyright Owned by —
        <strong
          >Lourde Matha Convent Metric Higher Secondary School,Udumalpet</strong
        >
        Open Source Icons From
        <a href="https://www.flaticon.com/" title="Flaticon"
          >www.flaticon.com</a
        >
        and <a href="https://icons8.com/" title="icons8">www.icons8.com</a>
      </v-card-text>
    </v-card>
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
export default {
  data: () => ({
    eventData: null,
    toggle: false,
    icons: [
      {
        icon: "mdi-youtube",
        link: "https://www.youtube.com/channel/UC3jmIA7i91OC4o4gboaAnXg"
      },
      {
        icon: "mdi-gmail",
        link: "mailto:lmcm.school@yahoo.com"
      },
      {
        icon: "mdi-login",
        link: "/admin"
      }
    ]
  }),
  components: {
    AppBar
  }
};
</script>

<style></style>
