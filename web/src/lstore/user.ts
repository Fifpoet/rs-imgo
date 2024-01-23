export function getUserCode(): number {
    let userCode = localStorage.getItem('userCode');

    // 如果userCode不存在，则生成一个6位随机整数并插入到localStorage中
    if (!userCode) {
        userCode = generateRandomInt(100000, 999999).toString();
        localStorage.setItem('userCode', userCode);
    }

    return parseInt(userCode);
}

function generateRandomInt(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}