import 'leaflet/dist/leaflet.css'
import L from 'leaflet'

const createMap = (divId, options) => {
    let map = L.map(divId, options)
    return map
}

//添加创建地图的函数
const createTileLayer = (map, url, options) => {
    let tileLayer = L.tileLayer(url, options)
    tileLayer.addTo(map)
    return tileLayer
}

export default {
    createMap,
    createTileLayer,
}