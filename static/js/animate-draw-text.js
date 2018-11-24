var ctx = document.getElementById("firstName").getContext("2d"),
    dashLen = 220, dashOffset = dashLen, speed = 6,
    txt = "Ryan James", x = 30, y=90, i = 0;

ctx.font = "50px cursive";
ctx.lineWidth = 1; ctx.lineJoin = "round"; ctx.globalAlpha = 2/3;
ctx.strokeStyle = ctx.fillStyle = "#FFFFFF";

(function loop() {
    if (txt[i] === ' ') {
        x = 30;
        y = 150;
    }
    ctx.clearRect(x, y, 60, 150);
    ctx.setLineDash([dashLen - dashOffset, dashOffset - speed]); // create a long dash mask
    dashOffset -= speed;                                         // reduce dash length
    ctx.strokeText(txt[i], x, y);                               // stroke letter

    if (dashOffset > 0) requestAnimationFrame(loop);             // animate
    else {
        ctx.fillText(txt[i], x, y);                               // fill final letter
        dashOffset = dashLen;                                      // prep next char
        x += ctx.measureText(txt[i++]).width + ctx.lineWidth * Math.random();
        ctx.setTransform(1, 0, 0, 1, 0, 3 * Math.random());        // random y-delta
        ctx.rotate(Math.random() * 0.005);                         // random rotation
        if (i < txt.length) requestAnimationFrame(loop);
    }
})();
