// yarn add i --dev imagemin imagemin-jpegtran imagemin-pngquant
// run via: yarn node /Users/sheldon.hull/git/github.com/sheldonhull/digital-garden/frontmatter-scripts/image-optimize.js "$(realpath .)" "notes/assets/images/copilot-color-codes.png"
const path = require('path');

const arguments = process.argv;

(async () => {
  if (arguments && arguments.length > 0) {
    const imagemin = (await import('imagemin')).default;
    const imageminJpegtran = (await import('imagemin-jpegtran')).default;
    const imageminPngquant = (await import('imagemin-pngquant')).default;

    const fileArg = arguments[3]; // The file path

    await imagemin([fileArg], {
      destination: path.dirname(fileArg),
      glob: false,
      plugins: [imageminJpegtran(), imageminPngquant()],
    });

    console.log(`Optimized image ${path.basename(fileArg)}`);
  }
})();
