function latLongToTileXY(latitude: number, longitude: number, levelOfDetail: number): [number, number] {
    const MinLatitude = -85.05112878;
    const MaxLatitude = 85.05112878;
    const MinLongitude = -180;
    const MaxLongitude = 180;

    latitude = clip(latitude, MinLatitude, MaxLatitude);
    longitude = clip(longitude, MinLongitude, MaxLongitude);

    const x = (longitude + 180) / 360;
    const sinLatitude = Math.sin((latitude * Math.PI) / 180);
    const y = 0.5 - Math.log((1 + sinLatitude) / (1 - sinLatitude)) / (4 * Math.PI);

    const size = mapSize(levelOfDetail);
    const tileX = clip(x * size + 0.5, 0, size - 1) / 256;
    const tileY = clip(y * size + 0.5, 0, size - 1) / 256;

    return [tileX, tileY];
}

function clip(value: number, min: number, max: number): number {
    return Math.min(Math.max(value, min), max);
}

function mapSize(levelOfDetail: number): number {
    return 256 * Math.pow(2, levelOfDetail);
}

export {
    latLongToTileXY
}