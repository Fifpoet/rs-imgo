<template>
  <div class="map-container" id="map-container"></div>
</template>

<script>
export default {
  name: "mapView",
  components: {},
  data() {
    return {
      map: null,
      OSMUrl: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    };
  },
  mounted() {
    this.map = this.$utils.map.createMap("map-container");
    // 设施地图视图 中心位置
    this.map.setView([30.35, 114.18], 13);

    // 加载 open street map和mapbox 图层服务
    this.$utils.map.createTileLayer(this.map, this.OSMUrl, {maxZoom: 19});
  }
};
</script>
<style scoped>
.map-container {
  position: absolute;
  left: 0;
  top: 100px;
  width: 100%;
  height: 100%;
}
</style>