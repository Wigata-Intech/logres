### Start Phoenix development server

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Run the Phoenix development server to start building with Tailwind CSS.

```bash
mix phx.server
```

--------------------------------

### justify-items-start Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-start` utility to justify grid items against the start of their inline axis.

```html
<div class="grid justify-items-start ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div>  <div>05</div>  <div>06</div></div>
```

--------------------------------

### Create a new TanStack Start project

Source: https://tailwindcss.com/docs/installation/framework-guides/tanstack-start

Initialize a new TanStack Start project using `create-start-app` and navigate into the project directory.

```bash
npx create-start-app@latest my-projectcd my-project
```

--------------------------------

### Start the Ember.js build process

Source: https://tailwindcss.com/docs/installation/framework-guides/emberjs

Run the `npm run start` command to build and serve the Ember.js application.

```bash
npm run start
```

--------------------------------

### justify-self-start Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-start` utility to align a grid item to the start of its inline axis.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-start ...">02</div>  <!-- ... --></div>
```

--------------------------------

### Start the build process

Source: https://tailwindcss.com/docs/installation/framework-guides/laravel/mix

Run the `npm run watch` command to start your build process, which will compile your assets and apply Tailwind CSS.

```bash
npm run watch
```

--------------------------------

### Basic transition example

Source: https://tailwindcss.com/docs/transition-property

Use utilities like `transition` and `transition-colors` to specify which properties should transition when they change. This example shows a button that changes background color and position on hover.

```html
<button class="bg-blue-500 transition delay-150 duration-300 ease-in-out hover:-translate-y-1 hover:scale-110 hover:bg-indigo-500 ...">  Save Changes</button>
```

--------------------------------

### Responsive Isolation Example

Source: https://tailwindcss.com/docs/isolation

Apply isolation utilities responsively by prefixing them with breakpoint variants like `md:`. This example applies `isolation-auto` at medium screen sizes and above.

```html
<div class="isolate md:isolation-auto ...">  <!-- ... --></div>
```

--------------------------------

### Snapping to the start

Source: https://tailwindcss.com/docs/scroll-snap-align

Use the `snap-start` utility to snap an element to its start when being scrolled inside a snap container.

```html
<div class="snap-x ...">  <div class="snap-start ...">    <img src="/img/vacation-01.jpg" />  </div>  <div class="snap-start ...">    <img src="/img/vacation-02.jpg" />  </div>  <div class="snap-start ...">    <img src="/img/vacation-03.jpg" />  </div>  <div class="snap-start ...">    <img src="/img/vacation-04.jpg" />  </div>  <div class="snap-start ...">    <img src="/img/vacation-05.jpg" />  </div>  <div class="snap-start ...">    <img src="/img/vacation-06.jpg" />  </div></div>
```

--------------------------------

### Start Gatsby development server

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Run `gatsby develop` to start the local development server and build your project.

```Terminal
gatsby develop
```

--------------------------------

### Install Webpack Encore and dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/symfony

Remove default asset handling bundles and install Webpack Encore with Turbo and Stimulus support.

```bash
composer remove symfony/ux-turbo symfony/asset-mapper symfony/stimulus-bundle
composer require symfony/webpack-encore-bundle symfony/ux-turbo symfony/stimulus-bundle
```

--------------------------------

### Install Tailwind CLI

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Download and install the standalone Tailwind CSS command-line interface.

```bash
mix tailwind.install
```

--------------------------------

### Basic mix-blend-mode Example

Source: https://tailwindcss.com/docs/mix-blend-mode

Use utilities like `mix-blend-overlay` and `mix-blend-soft-light` to control how an element's content and background is blended with other content in the same stacking context. This example uses `mix-blend-multiply`.

```html
<div class="flex justify-center -space-x-14">  <div class="bg-blue-500 mix-blend-multiply ..."></div>  <div class="bg-pink-500 mix-blend-multiply ..."></div></div>
```

--------------------------------

### Responsive break-after Example

Source: https://tailwindcss.com/docs/break-after

Shows how to apply break-after utilities responsively using breakpoint variants like `md:`. This example applies `break-after-column` by default and switches to `break-after-auto` at medium screen sizes and above.

```html
<div class="break-after-column md:break-after-auto ...">  <!-- ... --></div>
```

--------------------------------

### Basic transition-behavior Example

Source: https://tailwindcss.com/docs/transition-behavior

Use the `transition-discrete` utility to enable transitions for properties with discrete values, such as elements changing from `hidden` to `block`. This example shows interactive checkboxes controlling the visibility and transition behavior of buttons.

```html
<label class="peer ...">  <input type="checkbox" checked /></label><button class="hidden transition-all not-peer-has-checked:opacity-0 peer-has-checked:block ...">  I hide</button><label class="peer ...">  <input type="checkbox" checked /></label><button class="hidden transition-all transition-discrete not-peer-has-checked:opacity-0 peer-has-checked:block ...">  I fade out</button>
```

--------------------------------

### Responsive break-inside Example

Source: https://tailwindcss.com/docs/break-inside

Shows how to apply `break-inside` utilities responsively using Tailwind's breakpoint variants. This example applies `break-inside-avoid-column` at small sizes and `break-inside-auto` at medium screens and up.

```html
<div class="break-inside-avoid-column md:break-inside-auto ...">  <!-- ... --></div>
```

--------------------------------

### Create a new Laravel project

Source: https://tailwindcss.com/docs/installation/framework-guides/laravel/vite

Use the Laravel installer to create a new project and navigate into its directory.

```bash
laravel new my-projectcd my-project
```

--------------------------------

### Start Astro development server

Source: https://tailwindcss.com/docs/installation/framework-guides/astro

Run the development server to build and preview your Astro project with Tailwind CSS enabled.

```bash
npm run dev
```

--------------------------------

### Basic mask-composite examples

Source: https://tailwindcss.com/docs/mask-composite

Demonstrates the use of `mask-add`, `mask-subtract`, `mask-intersect`, and `mask-exclude` utilities to control mask combination. Each example applies different composite modes to an element with a background image and specific mask configurations.

```html
<div class="mask-add mask-[url(/img/circle.png),url(/img/circle.png)] mask-[position:30%_50%,70%_50%] bg-[url(/img/mountains.jpg)]"></div><div class="mask-subtract mask-[url(/img/circle.png),url(/img/circle.png)] mask-[position:30%_50%,70%_50%] bg-[url(/img/mountains.jpg)]"></div><div class="mask-intersect mask-[url(/img/circle.png),url(/img/circle.png)] mask-[position:30%_50%,70%_50%] bg-[url(/img/mountains.jpg)]"></div><div class="mask-exclude mask-[url(/img/circle.png),url(/img/circle.png)] mask-[position:30%_50%,70%_50%] bg-[url(/img/mountains.jpg)]"></div>
```

--------------------------------

### Start the Tailwind CSS build process

Source: https://tailwindcss.com/docs/installation/framework-guides/ruby-on-rails

Run the development server to compile Tailwind CSS.

```bash
./bin/dev
```

--------------------------------

### Create new Symfony project

Source: https://tailwindcss.com/docs/installation/framework-guides/symfony

Initialize a new Symfony web application project using the Symfony Installer.

```bash
symfony new --webapp my-project
cd my-project
```

--------------------------------

### Install Tailwind CSS and PostCSS plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/parcel

Install the required Tailwind CSS package and its peer dependencies via npm.

```bash
npm install tailwindcss @tailwindcss/postcss
```

--------------------------------

### Negative Scroll Padding Inline Start

Source: https://tailwindcss.com/docs/scroll-padding

Prefix scroll padding utilities with a dash to use negative values. For example, `-scroll-ps-6` for negative inline start padding.

```html
<div class="-scroll-ps-6 snap-x ...">  <!-- ... --></div>
```

--------------------------------

### HTML Example for transform-flat and transform-3d

Source: https://tailwindcss.com/docs/transform-style

This example demonstrates the visual difference between `transform-flat` (2D rendering) and `transform-3d` (3D rendering) for child elements within a parent container.

```html
<div class="size-20 transform-flat ...">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-x-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div><div class="size-20 transform-3d ...">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-x-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div>
```

--------------------------------

### Container Query CSS Example

Source: https://tailwindcss.com/docs/responsive-design

Example of a CSS rule using a container query with a minimum width of 16rem.

```css
@container (width >= 16rem) { … }
```

--------------------------------

### Install Tailwind CSS and PostCSS dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Install `@tailwindcss/postcss`, its peer dependencies, and `gatsby-plugin-postcss` using npm.

```Terminal
npm install @tailwindcss/postcss tailwindcss postcss gatsby-plugin-postcss
```

--------------------------------

### Install Tailwind CSS and PostCSS dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/emberjs

Install Tailwind CSS, its PostCSS plugin, PostCSS itself, and the PostCSS loader using npm.

```bash
npm install tailwindcss @tailwindcss/postcss postcss postcss-loader
```

--------------------------------

### Install Tailwind CSS and dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/nextjs

Install the required npm packages: @tailwindcss/postcss, tailwindcss, and postcss.

```bash
npm install tailwindcss @tailwindcss/postcss postcss
```

--------------------------------

### Custom Scroll Padding Inline Start

Source: https://tailwindcss.com/docs/scroll-padding

Use arbitrary value syntax like `scroll-pl-[<value>]` for custom scroll padding. For example, `scroll-pl-[24rem]`.

```html
<div class="scroll-pl-[24rem] ...">  <!-- ... --></div>
```

--------------------------------

### Create Parcel project structure

Source: https://tailwindcss.com/docs/installation/framework-guides/parcel

Initialize a new Parcel project with basic directory and file setup.

```bash
mkdir my-projectcd my-projectnpm init -ynpm install parcelmkdir srctouch src/index.html
```

--------------------------------

### Start Parcel development server

Source: https://tailwindcss.com/docs/installation/framework-guides/parcel

Run the Parcel build process to serve your project during development.

```bash
npx parcel src/index.html
```

--------------------------------

### Install Tailwind CSS dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/meteor

Install Tailwind CSS, the PostCSS plugin, and required peer dependencies via npm.

```bash
npm install tailwindcss @tailwindcss/postcss postcss postcss-load-config
```

--------------------------------

### Basic scroll-padding example

Source: https://tailwindcss.com/docs/scroll-padding

Use scroll-pt, scroll-pr, scroll-pb, and scroll-pl utilities to set the scroll offset within a snap container. This example demonstrates setting left and top padding.

```html
<div class="snap-x scroll-pl-6 ...">
  <div class="snap-start ...">
    <img src="/img/vacation-01.jpg" />
  </div>
  <div class="snap-start ...">
    <img src="/img/vacation-02.jpg" />
  </div>
  <div class="snap-start ...">
    <img src="/img/vacation-03.jpg" />
  </div>
  <div class="snap-start ...">
    <img src="/img/vacation-04.jpg" />
  </div>
  <div class="snap-start ...">
    <img src="/img/vacation-05.jpg" />
  </div>
</div>
```

--------------------------------

### Install Tailwind CSS and Vite plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/astro

Install the necessary Tailwind CSS packages, including the Vite plugin, as project dependencies.

```bash
npm install tailwindcss @tailwindcss/vite
```

--------------------------------

### Install Tailwind CSS and PostCSS dependencies

Source: https://tailwindcss.com/docs/installation/framework-guides/angular

Install Tailwind CSS, the PostCSS plugin, and PostCSS itself via npm with the --force flag.

```bash
npm install tailwindcss @tailwindcss/postcss postcss --force
```

--------------------------------

### justify-items-center-safe Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-center-safe` utility to justify grid items against the end of their inline axis. When there is not enough space available, it will align items to the start of the container instead of the center.

```html
<div class="grid grid-flow-col justify-items-center-safe ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### justify-self-end-safe Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-end-safe` utility to align a grid item to the end of its inline axis. When there is not enough space available, it will align the item to the start of the container instead of the end.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-end-safe ...">02</div>  <!-- ... --></div>
```

--------------------------------

### justify-self-center-safe Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-center-safe` utility to align a grid item along the center of its inline axis. When there is not enough space available, it will align the item to the start of the container instead of the end.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-center-safe ...">02</div>  <!-- ... --></div>
```

--------------------------------

### Create a new Qwik project

Source: https://tailwindcss.com/docs/installation/framework-guides/qwik

Initialize a new Qwik project using `npm create qwik` if you don't have one set up already.

```bash
npm create qwik@latest empty my-projectcd my-project
```

--------------------------------

### Block Start Border Colors (Predefined)

Source: https://tailwindcss.com/docs/border-color

Use these utilities to set the `border-block-start-color` property with predefined color values from the Tailwind CSS palette. Examples include olive and mist color scales.

```html
<!-- Olive -->
<div class="border-bs-olive-950 ..."></div>

<!-- Mist -->
<div class="border-bs-mist-50 ..."></div>
<div class="border-bs-mist-100 ..."></div>
<div class="border-bs-mist-200 ..."></div>
<div class="border-bs-mist-300 ..."></div>
<div class="border-bs-mist-400 ..."></div>
<div class="border-bs-mist-500 ..."></div>
<div class="border-bs-mist-600 ..."></div>
<div class="border-bs-mist-700 ..."></div>
<div class="border-bs-mist-800 ..."></div>
<div class="border-bs-mist-900 ..."></div>
<div class="border-bs-mist-950 ..."></div>

<!-- Taupe -->
<div class="border-bs-taupe-50 ..."></div>
<div class="border-bs-taupe-100 ..."></div>
<div class="border-bs-taupe-200 ..."></div>
<div class="border-bs-taupe-300 ..."></div>
<div class="border-bs-taupe-400 ..."></div>
<div class="border-bs-taupe-500 ..."></div>
<div class="border-bs-taupe-600 ..."></div>
<div class="border-bs-taupe-700 ..."></div>
<div class="border-bs-taupe-800 ..."></div>
<div class="border-bs-taupe-900 ..."></div>
<div class="border-bs-taupe-950 ..."></div>
```

--------------------------------

### Start Angular development server

Source: https://tailwindcss.com/docs/installation/framework-guides/angular

Run the Angular development server to build and serve your project with Tailwind CSS enabled.

```bash
ng serve
```

--------------------------------

### justify-items-end-safe Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-end-safe` utility to justify grid items against the end of their inline axis. When there is not enough space available, it will align items to the start of the container instead of the end.

```html
<div class="grid grid-flow-col justify-items-end-safe ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Align Content Start in Tailwind CSS

Source: https://tailwindcss.com/docs/align-content

Use `content-start` to pack rows in a container against the start of the cross axis in a grid layout.

```html
<div class="grid h-56 grid-cols-3 content-start gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div>  <div>05</div></div>
```

--------------------------------

### Create a new SvelteKit project

Source: https://tailwindcss.com/docs/installation/framework-guides/sveltekit

Initializes a new SvelteKit project and navigates into its directory using `npx sv create`.

```bash
npx sv create my-projectcd my-project
```

--------------------------------

### Install Tailwind CSS CLI via npm

Source: https://tailwindcss.com/docs/installation/tailwind-cli

Use npm to install the 'tailwindcss' and '@tailwindcss/cli' packages for project development.

```Terminal
npm install tailwindcss @tailwindcss/cli
```

--------------------------------

### Install tailwindcss-rails gem

Source: https://tailwindcss.com/docs/installation/framework-guides/ruby-on-rails

Add the tailwindcss-rails gem to your project and run the installation command.

```bash
bundle add tailwindcss-rails./bin/rails tailwindcss:install
```

--------------------------------

### Add Play CDN script to HTML

Source: https://tailwindcss.com/docs/installation/play-cdn

Include the Tailwind CSS Play CDN script tag in the HTML head to enable Tailwind utility classes. This is the minimal setup required to start using Tailwind without any build step.

```html
<!doctype html><html>  <head>    <meta charset="UTF-8" />    <meta name="viewport" content="width=device-width, initial-scale=1.0" />    <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>  </head>  <body>    <h1 class="text-3xl font-bold underline">      Hello world!    </h1>  </body></html>
```

--------------------------------

### Create a new Astro project

Source: https://tailwindcss.com/docs/installation/framework-guides/astro

Initialize a new Astro project using `create astro` and navigate into the project directory.

```bash
npm create astro@latest my-projectcd my-project
```

--------------------------------

### Responsive Break Before Example

Source: https://tailwindcss.com/docs/break-before

Prefix a `break-before` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="break-before-column md:break-before-auto ...">  <!-- ... --></div>
```

--------------------------------

### Place grid items at the start

Source: https://tailwindcss.com/docs/place-items

Use `place-items-start` to align grid items to the start of their grid areas on both axes. This is useful for left-aligning content within grid cells.

```html
<div class="grid grid-cols-3 place-items-start gap-4 ...">
  <div>01</div>
  <div>02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### Setting Inline Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-inline-{color}-{shade}` utility to set the color of the inline start and end borders. This is useful for logical properties in different writing modes. For example, `border-inline-red-500` sets the inline borders to a medium red.

```html
<div class="border-inline-red-500">...</div>
```

--------------------------------

### justify-start

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-start` utility to justify items against the start of the container's main axis.

```html
<div class="flex justify-start ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Basic Tab Size Example

Source: https://tailwindcss.com/docs/tab-size

Use `tab-<number>` utilities like `tab-2` and `tab-8` to control the size of tab characters. This example shows the JavaScript function with different tab sizes.

```javascript
function indent() {	return 'tabbed';}
```

```javascript
function indent() {	return 'tabbed';}
```

--------------------------------

### Create a new Nuxt project

Source: https://tailwindcss.com/docs/installation/framework-guides/nuxt

Initialize a new Nuxt project using `create-nuxt` and navigate into the project directory.

```bash
npm create nuxt my-project
cd my-project
```

--------------------------------

### Responsive place-items

Source: https://tailwindcss.com/docs/place-items

Apply `place-items` utilities conditionally at different screen sizes using responsive variants. This example applies `place-items-start` by default and switches to `place-items-center` at medium screen sizes and above.

```html
<div class="grid place-items-start md:place-items-center ...">
  <!-- ... -->
</div>
```

--------------------------------

### Create React Router project

Source: https://tailwindcss.com/docs/installation/framework-guides/react-router

Initialize a new React Router project using Create React Router.

```bash
npx create-react-router@latest my-projectcd my-project
```

--------------------------------

### Create a new Vite project

Source: https://tailwindcss.com/docs

Initialize a new Vite project using Create Vite and navigate into the project directory.

```bash
npm create vite@latest my-projectcd my-project
```

--------------------------------

### Setting Block Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-block-{color}-{shade}` utility to set the color of the block start and end borders. This is useful for logical properties in different writing modes. For example, `border-block-blue-800` sets the block borders to a dark blue.

```html
<div class="border-block-blue-800">...</div>
```

--------------------------------

### Percentage-based max-block-size Utilities

Source: https://tailwindcss.com/docs/max-block-size

Employ `max-block-full` or `max-block-<fraction>` utilities for percentage-based maximum block sizes. Examples include `max-block-9/10`, `max-block-3/4`, `max-block-1/2`, and `max-block-full`.

```html
<div class="block-96 ...">  <div class="block-full max-block-9/10 ...">max-block-9/10</div>  <div class="block-full max-block-3/4 ...">max-block-3/4</div>  <div class="block-full max-block-1/2 ...">max-block-1/2</div>  <div class="block-full max-block-full ...">max-block-full</div></div>
```

--------------------------------

### Styling Elements on Initial Render with @starting-style

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `starting` variant to define the initial appearance of an element when it's first rendered or transitions from `display: none`. This is useful for animations and transitions.

```html
<div>  <button popovertarget="my-popover">Check for updates</button>  <div popover id="my-popover" class="opacity-0 starting:open:opacity-0 ...">    <!-- ... -->  </div></div>
```

--------------------------------

### Responsive Profile Card with State Styles

Source: https://tailwindcss.com/docs/styling-with-utility-classes

This example showcases a responsive profile card built entirely with utility classes. It includes styling for different screen sizes, hover effects on the button, and active states.

```html
<div class="flex flex-col gap-2 p-8 sm:flex-row sm:items-center sm:gap-6 sm:py-4 ...">
  <img class="mx-auto block h-24 rounded-full sm:mx-0 sm:shrink-0" src="/img/erin-lindford.jpg" alt="" />
  <div class="space-y-2 text-center sm:text-left">
    <div class="space-y-0.5">
      <p class="text-lg font-semibold text-black">Erin Lindford</p>
      <p class="font-medium text-gray-500">Product Engineer</p>
    </div>
    <button class="border-purple-200 text-purple-600 hover:border-transparent hover:bg-purple-600 hover:text-white active:bg-purple-700 ...">
      Message
    </button>
  </div>
</div>
```

--------------------------------

### Create a new Phoenix project

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Initialize a new Phoenix project using the mix command.

```bash
mix phx.new myprojectcd myproject
```

--------------------------------

### Basic break-after Example

Source: https://tailwindcss.com/docs/break-after

Demonstrates how to use break-after utilities like `break-after-column` and `break-after-page` to control element breaking behavior within a two-column layout.

```html
<div class="columns-2">  <p>Well, let me tell you something, ...</p>  <p class="break-after-column">Sure, go ahead, laugh...</p>  <p>Maybe we can live without...</p>  <p>Look. If you think this is...</p></div>
```

--------------------------------

### Basic Isolation Example

Source: https://tailwindcss.com/docs/isolation

Use the `isolate` utility to explicitly create a new stacking context for an element. Use `isolation-auto` to allow the browser to determine stacking context creation.

```html
<div class="isolate ...">  <!-- ... --></div>
```

--------------------------------

### Create Meteor project

Source: https://tailwindcss.com/docs/installation/framework-guides/meteor

Initialize a new Meteor project using the CLI and navigate to the project directory.

```bash
npx meteor create my-projectcd my-project
```

--------------------------------

### Basic grid-auto-flow Example

Source: https://tailwindcss.com/docs/grid-auto-flow

Demonstrates how to use `grid-flow-row-dense` along with `grid-cols-3` and `grid-rows-3` to control the auto-placement of grid items. Useful for creating complex grid layouts where item size and order matter.

```html
<div class="grid grid-flow-row-dense grid-cols-3 grid-rows-3 ...">
  <div class="col-span-2">01</div>
  <div class="col-span-2">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
</div>
```

--------------------------------

### CSS accent-color property for accent-olive-600

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-600` utility.

```css
accent-color: var(--color-olive-600); /* oklch(46.6% 0.025 107.3) */
```

--------------------------------

### Responsive grid-auto-rows with breakpoints

Source: https://tailwindcss.com/docs/grid-auto-rows

Demonstrates applying `grid-auto-rows` utilities responsively using breakpoint prefixes like `md:`. This example switches from `auto-rows-max` to `auto-rows-min` at the medium breakpoint.

```html
<div class="grid grid-flow-row auto-rows-max md:auto-rows-min ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-olive-700

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-700` utility.

```css
accent-color: var(--color-olive-700); /* oklch(39.4% 0.023 107.4) */
```

--------------------------------

### Flexbox align-self: start

Source: https://tailwindcss.com/docs/align-self

Use the `self-start` utility to align an item to the start of the container's cross axis, overriding the container's `align-items` value.

```html
<div class="flex items-stretch ...">  <div>01</div>  <div class="self-start ...">02</div>  <div>03</div></div>
```

--------------------------------

### Applying Responsive Utility Variants

Source: https://tailwindcss.com/docs/responsive-design

Prefix utility classes with breakpoint names (e.g., `md:`, `lg:`) to apply them only at specific screen sizes and larger. This example shows width adjustments at different breakpoints.

```html
<!-- Width of 16 by default, 32 on medium screens, and 48 on large screens -->
<img class="w-16 md:w-32 lg:w-48" src="..." />
```

--------------------------------

### CSS accent-color property for accent-olive-100

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-100` utility.

```css
accent-color: var(--color-olive-100); /* oklch(96.6% 0.005 106.5) */
```

--------------------------------

### Basic Aspect Ratio

Source: https://tailwindcss.com/docs/aspect-ratio

Use `aspect-<ratio>` utilities like `aspect-3/2` to give an element a specific aspect ratio. Resize the example to see the expected behavior.

```html
<img class="aspect-3/2 object-cover ..." src="/img/villas.jpg" />
```

--------------------------------

### CSS accent-color property for accent-mauve-700

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-700` utility.

```css
accent-color: var(--color-mauve-700); /* oklch(36.4% 0.029 323.89) */
```

--------------------------------

### place-self-start alignment in grid

Source: https://tailwindcss.com/docs/place-self

Align a grid item to the start on both axes using the place-self-start class.

```html
<div class="grid grid-cols-3 gap-4 ...">
  <div>01</div>
  <div class="place-self-start ...">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### justify-items-stretch Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-stretch` utility to stretch items along their inline axis.

```html
<div class="grid justify-items-stretch ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div>  <div>05</div>  <div>06</div></div>
```

--------------------------------

### CSS accent-color property for accent-olive-400

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-400` utility.

```css
accent-color: var(--color-olive-400); /* oklch(73.7% 0.021 106.9) */
```

--------------------------------

### Responsive Perspective

Source: https://tailwindcss.com/docs/perspective

Apply perspective utilities at specific breakpoints using responsive prefixes like `md:`. This example applies `perspective-midrange` by default and `perspective-dramatic` at medium screens and above.

```html
<div class="perspective-midrange md:perspective-dramatic ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-mauve-600

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-600` utility.

```css
accent-color: var(--color-mauve-600); /* oklch(43.5% 0.029 321.78) */
```

--------------------------------

### CSS accent-color property for accent-mauve-400

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-400` utility.

```css
accent-color: var(--color-mauve-400); /* oklch(71.1% 0.019 323.02) */
```

--------------------------------

### CSS accent-color property for accent-olive-500

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-500` utility.

```css
accent-color: var(--color-olive-500); /* oklch(58% 0.031 107.3) */
```

--------------------------------

### Basic grid-auto-rows example

Source: https://tailwindcss.com/docs/grid-auto-rows

Demonstrates using `auto-rows-max` to control implicitly created grid rows. This sets the row size to `max-content`.

```html
<div class="grid grid-flow-row auto-rows-max">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### CSS accent-color property for accent-mauve-900

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-900` utility.

```css
accent-color: var(--color-mauve-900); /* oklch(21.2% 0.019 322.12) */
```

--------------------------------

### Basic break-inside Example

Source: https://tailwindcss.com/docs/break-inside

Demonstrates how to use `break-inside-avoid-column` to prevent an element from breaking across columns. This is useful for keeping related content together.

```html
<div class="columns-2">  <p>Well, let me tell you something, ...</p>  <p class="break-inside-avoid-column">Sure, go ahead, laugh...</p>  <p>Maybe we can live without...</p>  <p>Look. If you think this is...</p></div>
```

--------------------------------

### CSS accent-color property for accent-olive-300

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-300` utility.

```css
accent-color: var(--color-olive-300); /* oklch(88% 0.011 106.6) */
```

--------------------------------

### Example HTML with Tailwind CSS classes

Source: https://tailwindcss.com/docs/installation

Demonstrates how to include the compiled CSS in your HTML and apply Tailwind's utility classes to style elements.

```HTML
<!doctype html><html><head>  <meta charset="UTF-8">  <meta name="viewport" content="width=device-width, initial-scale=1.0">  <link href="/src/style.css" rel="stylesheet"></head><body>  <h1 class="text-3xl font-bold underline">    Hello world!  </h1></body></html>
```

--------------------------------

### CSS accent-color property for accent-mauve-100

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-100` utility.

```css
accent-color: var(--color-mauve-100); /* oklch(96% 0.003 325.6) */
```

--------------------------------

### Create new Angular project with Angular CLI

Source: https://tailwindcss.com/docs/installation/framework-guides/angular

Initialize a new Angular project with CSS as the stylesheet format using the Angular CLI.

```bash
ng new my-project --style css
cd my-project
```

--------------------------------

### CSS accent-color property for accent-mauve-800

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-800` utility.

```css
accent-color: var(--color-mauve-800); /* oklch(26.3% 0.024 320.12) */
```

--------------------------------

### CSS accent-color property for accent-olive-50

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-50` utility.

```css
accent-color: var(--color-olive-50); /* oklch(98.8% 0.003 106.5) */
```

--------------------------------

### Create a new Gatsby project

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Use the Gatsby CLI to initialize a new project and navigate into its directory.

```Terminal
gatsby new my-projectcd my-project
```

--------------------------------

### CSS accent-color property for accent-olive-200

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-olive-200` utility.

```css
accent-color: var(--color-olive-200); /* oklch(93% 0.007 106.5) */
```

--------------------------------

### Set Border-inline-start Color with Red Palette

Source: https://tailwindcss.com/docs/border-color

This example shows how to set the inline-start border color using shades from the red color palette.

```css
border-inline-start-color: var(--color-red-950); /* oklch(25.8% 0.092 26.042) */
```

--------------------------------

### Align items to the start

Source: https://tailwindcss.com/docs/align-items

Use the `items-start` utility to align flex items to the beginning of the container's cross axis.

```html
<div class="flex items-start ...">
  <div class="py-4">01</div>
  <div class="py-12">02</div>
  <div class="py-8">03</div>
</div>
```

--------------------------------

### Create a new Rails project

Source: https://tailwindcss.com/docs/installation/framework-guides/ruby-on-rails

Use the Rails Command Line to create a new project and navigate into its directory.

```bash
rails new my-projectcd my-project
```

--------------------------------

### CSS accent-color property for accent-mauve-50

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-50` utility.

```css
accent-color: var(--color-mauve-50); /* oklch(98.5% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-mauve-950

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-950` utility.

```css
accent-color: var(--color-mauve-950); /* oklch(14.5% 0.008 326) */
```

--------------------------------

### CSS accent-color property for accent-mauve-500

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-500` utility.

```css
accent-color: var(--color-mauve-500); /* oklch(54.2% 0.034 322.5) */
```

--------------------------------

### CSS accent-color property for accent-mauve-200

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-200` utility.

```css
accent-color: var(--color-mauve-200); /* oklch(92.2% 0.005 325.62) */
```

--------------------------------

### CSS accent-color property for accent-mauve-300

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-mauve-300` utility.

```css
accent-color: var(--color-mauve-300); /* oklch(86.5% 0.012 325.68) */
```

--------------------------------

### Align content to start in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-start` to align grid items to the beginning of both the inline and block axes.

```html
<div class="grid h-48 grid-cols-2 place-content-start gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### CSS accent-color property for accent-stone-600

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-600` utility.

```css
accent-color: var(--color-stone-600); /* oklch(44.4% 0.011 73.639) */
```

--------------------------------

### Responsive Column Design

Source: https://tailwindcss.com/docs/columns

Apply column utilities responsively using breakpoint prefixes like `sm:`. This example applies `columns-2` and `gap-4` by default, switching to `columns-3` and `gap-8` at the `sm` breakpoint and above.

```html
<div class="columns-2 gap-4 sm:columns-3 sm:gap-8 ...">
  <img class="aspect-3/2 ..." src="/img/mountains-1.jpg" />
  <img class="aspect-square ..." src="/img/mountains-2.jpg" />
  <img class="aspect-square ..." src="/img/mountains-3.jpg" />
  <!-- ... -->
</div>
```

--------------------------------

### Responsive Resizing with Breakpoints

Source: https://tailwindcss.com/docs/resize

Apply resize utilities conditionally based on screen size using responsive prefixes. This example shows an element that is not resizable by default but becomes resizable at medium screen sizes and above.

```html
<div class="resize-none md:resize ...">  <!-- ... --></div>
```

--------------------------------

### Responsive line clamping

Source: https://tailwindcss.com/docs/line-clamp

Apply `line-clamp` utilities with responsive prefixes like `md:` to control text clamping at different screen sizes. This example applies 3 lines by default and 4 lines on medium screens and up.

```html
<div class="line-clamp-3 md:line-clamp-4 ...">
  <!-- ... -->
</div>
```

--------------------------------

### CSS accent-color property for accent-stone-700

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-700` utility.

```css
accent-color: var(--color-stone-700); /* oklch(37.4% 0.01 67.558) */
```

--------------------------------

### CSS accent-color property for accent-stone-400

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-400` utility.

```css
accent-color: var(--color-stone-400); /* oklch(70.9% 0.01 56.259) */
```

--------------------------------

### Using Logical Properties for Positioning

Source: https://tailwindcss.com/docs/top-right-bottom-left

Utilize `inset-s` and `inset-e` for inline start and end positioning, which adapt to text direction (LTR/RTL).

```html
<div dir="ltr">  <div class="relative size-32 ...">    <div class="absolute inset-s-0 top-0 size-14 ..."></div>  </div>  <div>    <div dir="rtl">      <div class="relative size-32 ...">        <div class="absolute inset-s-0 top-0 size-14 ..."></div>      </div>      <div></div>    </div>  </div></div>
```

--------------------------------

### Responsive transition-delay

Source: https://tailwindcss.com/docs/transition-delay

Prefix a `transition-delay` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example applies a delay of 150ms by default and 300ms on medium screens and up.

```html
<button class="delay-150 md:delay-300 ...">  <!-- ... --></button>
```

--------------------------------

### Masking Left Edge with Start and End

Source: https://tailwindcss.com/docs/mask-image

Apply a linear gradient mask to the left edge, specifying both the start and end points of the gradient using `mask-l-from-<value>` and `mask-l-to-<value>`. The background image is also applied.

```html
<div class="mask-l-from-50% mask-l-to-90% bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Masking Bottom Edge with Start and End

Source: https://tailwindcss.com/docs/mask-image

Apply a linear gradient mask to the bottom edge, specifying both the start and end points of the gradient using `mask-b-from-<value>` and `mask-b-to-<value>`. The background image is also applied.

```html
<div class="mask-b-from-20% mask-b-to-80% bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Basic Card Component with Utility Classes

Source: https://tailwindcss.com/docs/styling-with-utility-classes

This example demonstrates styling a card component using a variety of Tailwind utility classes for layout, spacing, appearance, and responsiveness. It includes dark mode variations.

```html
<div class="mx-auto flex max-w-sm items-center gap-x-4 rounded-xl bg-white p-6 shadow-lg outline outline-black/5 dark:bg-slate-800 dark:shadow-none dark:-outline-offset-1 dark:outline-white/10">
  <img class="size-12 shrink-0" src="/img/logo.svg" alt="ChitChat Logo" />
  <div>
    <div class="text-xl font-medium text-black dark:text-white">ChitChat</div>
    <p class="text-gray-500 dark:text-gray-400">You have a new message!</p>
  </div>
</div>
```

--------------------------------

### CSS accent-color property for accent-stone-900

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-900` utility.

```css
accent-color: var(--color-stone-900); /* oklch(21.6% 0.006 56.043) */
```

--------------------------------

### Basic Container Query Example

Source: https://tailwindcss.com/docs/responsive-design

Mark an element as an inline-size container using the `@container` class. Style child elements based on the container's size using variants like `@sm` and `@md`.

```html
<div class="@container">  <div class="flex flex-col @md:flex-row">    <!-- ... -->  </div></div>
```

--------------------------------

### Basic Perspective Example

Source: https://tailwindcss.com/docs/perspective

Demonstrates using `perspective-dramatic` and `perspective-normal` to control the z-plane distance. Each div contains nested elements with various transformations.

```html
<div class="size-20 perspective-dramatic ...">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-x-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div><div class="size-20 perspective-normal ...">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-x-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div>
```

--------------------------------

### CSS accent-color property for accent-stone-100

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-100` utility.

```css
accent-color: var(--color-stone-100); /* oklch(97% 0.001 106.424) */
```

--------------------------------

### Basic line-clamp Example

Source: https://tailwindcss.com/docs/line-clamp

Use line-clamp-3 to truncate multi-line text after three lines. This is useful for displaying previews of longer content.

```html
<article>
  <time>Mar 10, 2020</time>
  <h2>Boost your conversion rate</h2>
  <p class="line-clamp-3">
    Nulla dolor velit adipisicing duis excepteur esse in duis nostrud occaecat mollit incididunt deserunt sunt. Ut ut
    sunt laborum ex occaecat eu tempor labore enim adipisicing minim ad. Est in quis eu dolore occaecat excepteur fugiat
    dolore nisi aliqua fugiat enim ut cillum. Labore enim duis nostrud eu. Est ut eiusmod consequat irure quis deserunt
    ex. Enim laboris dolor magna pariatur. Dolor et ad sint voluptate sunt elit mollit officia ad enim sit consectetur
    enim.
  </p>
  <div>
    <img src="/img/lindsay.jpg" />
    Lindsay Walton
  </div>
</article>
```

--------------------------------

### Responsive transition property

Source: https://tailwindcss.com/docs/transition-property

Prefix a `transition-property` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example applies `transition-none` by default and `transition-all` from the medium breakpoint upwards.

```html
<button class="transition-none md:transition-all ...">  <!-- ... --></button>
```

--------------------------------

### Viewport Meta Tag Setup

Source: https://tailwindcss.com/docs/responsive-design

Include this meta tag in the `<head>` of your HTML document to ensure proper rendering of responsive designs.

```html
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
```

--------------------------------

### Create a new AdonisJS project

Source: https://tailwindcss.com/docs/installation/framework-guides/adonisjs

Initialize a new AdonisJS project using the `create adonisjs` command with the web kit.

```Terminal
npm init adonisjs@latest my-project -- --kit=webcd my-project
```

--------------------------------

### CSS accent-color property for accent-stone-800

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-800` utility.

```css
accent-color: var(--color-stone-800); /* oklch(26.8% 0.007 34.298) */
```

--------------------------------

### Responsive Marketing Page Component

Source: https://tailwindcss.com/docs/responsive-design

This example demonstrates a responsive component layout that adapts from a stacked view on small screens to a side-by-side layout on medium screens and larger. It utilizes flexbox and image sizing utilities conditionally.

```html
<div class="mx-auto max-w-md overflow-hidden rounded-xl bg-white shadow-md md:max-w-2xl">
  <div class="md:flex">
    <div class="md:shrink-0">
      <img
        class="h-48 w-full object-cover md:h-full md:w-48"
        src="/img/building.jpg"
        alt="Modern building architecture"
      />
    </div>
    <div class="p-8">
      <div class="text-sm font-semibold tracking-wide text-indigo-500 uppercase">Company retreats</div>
      <a href="#" class="mt-1 block text-lg leading-tight font-medium text-black hover:underline">
        Incredible accommodation for your team
      </a>
      <p class="mt-2 text-gray-500">
        Looking to take your team away on a retreat to enjoy awesome food and take in some sunshine? We have a list of
        places to do just that.
      </p>
    </div>
  </div>
</div>
```

--------------------------------

### CSS accent-color property for accent-stone-500

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-500` utility.

```css
accent-color: var(--color-stone-500); /* oklch(55.3% 0.013 58.071) */
```

--------------------------------

### Start Tailwind CLI build process

Source: https://tailwindcss.com/docs/installation/tailwind-cli

Run the Tailwind CLI to process your input CSS, scan for utility classes, and generate the final output CSS file, with '--watch' for continuous compilation.

```Terminal
npx @tailwindcss/cli -i ./src/input.css -o ./src/output.css --watch
```

--------------------------------

### CSS accent-color property for accent-stone-200

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-200` utility.

```css
accent-color: var(--color-stone-200); /* oklch(92.3% 0.003 48.717) */
```

--------------------------------

### CSS accent-color property for accent-stone-50

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-50` utility.

```css
accent-color: var(--color-stone-50); /* oklch(98.5% 0.001 106.423) */
```

--------------------------------

### Responsive min-height design

Source: https://tailwindcss.com/docs/min-height

Apply min-height utilities responsively by prefixing them with a breakpoint variant like md:. This example sets a default min-h-0 and applies min-h-full at medium screen sizes and above.

```html
<div class="h-24 min-h-0 md:min-h-full ...">  <!-- ... --></div>
```

--------------------------------

### Avatar Display within a Loop (Svelte Example)

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Renders a list of contributor avatars dynamically using a loop, demonstrating how to avoid actual markup duplication in a real project.

```html
<div>  <div class="flex items-center space-x-2 text-base">
    <h4 class="font-semibold text-slate-900">Contributors</h4>
    <span class="bg-slate-100 px-2 py-1 text-xs font-semibold text-slate-700 ...">204</span>
  </div>
  <div class="mt-3 flex -space-x-2 overflow-hidden">
    {#each contributors as user}
      <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src={user.avatarUrl} alt={user.handle} />
    {/each}
  </div>
  <div class="mt-3 text-sm font-medium">
    <a href="#" class="text-blue-500">+ 198 others</a>
  </div>
</div>
```

--------------------------------

### Apply block start padding with pbs-<number>

Source: https://tailwindcss.com/docs/padding

Use `pbs-<number>` utilities to set the `padding-block-start` logical property, which maps to either the top or bottom side based on the writing mode.

```html
<div class="pbs-8 ...">pbs-8</div>
```

--------------------------------

### Basic min-inline-size utilities

Source: https://tailwindcss.com/docs/min-inline-size

Use `min-inline-<number>` utilities to set a fixed minimum inline size based on the spacing scale. Examples include `min-inline-80` through `min-inline-24`.

```html
<div class="inline-20 ...">  <div class="min-inline-80 ...">min-inline-80</div>  <div class="min-inline-64 ...">min-inline-64</div>  <div class="min-inline-48 ...">min-inline-48</div>  <div class="min-inline-40 ...">min-inline-40</div>  <div class="min-inline-32 ...">min-inline-32</div>  <div class="min-inline-24 ...">min-inline-24</div></div>
```

--------------------------------

### CSS accent-color property for accent-stone-300

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-300` utility.

```css
accent-color: var(--color-stone-300); /* oklch(86.9% 0.005 56.366) */
```

--------------------------------

### Responsive grid-auto-flow

Source: https://tailwindcss.com/docs/grid-auto-flow

Shows how to apply `grid-auto-flow` utilities responsively using breakpoint prefixes like `md:`. This allows the grid's auto-placement behavior to change at different screen sizes.

```html
<div class="grid grid-flow-col md:grid-flow-row ...">
  <!-- ... -->
</div>
```

--------------------------------

### mask-auto utility

Source: https://tailwindcss.com/docs/mask-size

Use the `mask-auto` utility to display the mask image at its default size. This example also sets a background image.

```html
<div class="mask-auto mask-[url(/img/scribble.png)] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### CSS accent-color property for accent-stone-950

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-stone-950` utility.

```css
accent-color: var(--color-stone-950); /* oklch(14.7% 0.004 49.25) */
```

--------------------------------

### Responsive grid-auto-columns Design

Source: https://tailwindcss.com/docs/grid-auto-columns

Demonstrates applying `grid-auto-columns` utilities responsively using breakpoint variants like `md:`. This allows column sizing to adapt based on screen size, for example, switching from `auto-cols-max` to `auto-cols-min` at medium breakpoints.

```html
<div class="grid grid-flow-col auto-cols-max md:auto-cols-min ...">  <!-- ... --></div>
```

--------------------------------

### place-self responsive design with breakpoint

Source: https://tailwindcss.com/docs/place-self

Apply place-self utilities conditionally at different screen sizes using breakpoint prefixes like md:.

```html
<div class="place-self-start md:place-self-end ...">
  <!-- ... -->
</div>
```

--------------------------------

### Create a SolidJS Project

Source: https://tailwindcss.com/docs/installation/framework-guides/solidjs

Use npx degit to scaffold a new SolidJS project with the Vite template.

```bash
npx degit solidjs/templates/js my-projectcd my-project
```

--------------------------------

### Responsive Background Blend Mode

Source: https://tailwindcss.com/docs/background-blend-mode

Apply background blend mode utilities at specific breakpoints using responsive prefixes. This example applies 'lighten' by default and 'darken' on medium screens and up.

```html
<div class="bg-blue-500 bg-[url(/img/mountains.jpg)] bg-blend-lighten md:bg-blend-darken ...">  <!-- ... --></div>
```

--------------------------------

### Compare wrap-break-word and wrap-anywhere behavior

Source: https://tailwindcss.com/docs/overflow-wrap

This example demonstrates the difference between `wrap-break-word` and `wrap-anywhere` when applied to text within a flex container, showing how `wrap-anywhere` considers mid-word breaks for intrinsic sizing.

```html
<div class="flex max-w-sm">  <img class="size-16 rounded-full" src="/img/profile.jpg" />  <div class="wrap-break-word">    <p class="font-medium">Jay Riemenschneider</p>    <p>jason.riemenschneider@vandelayindustries.com</p>  </div></div><div class="flex max-w-sm">  <img class="size-16 rounded-full" src="/img/profile.jpg" />  <div class="wrap-anywhere">    <p class="font-medium">Jay Riemenschneider</p>    <p>jason.riemenschneider@vandelayindustries.com</p>  </div></div>
```

--------------------------------

### Responsive mask-composite utility

Source: https://tailwindcss.com/docs/mask-composite

Shows how to apply `mask-composite` utilities conditionally at different screen sizes using responsive variants like `md:`. This example applies `mask-add` by default and switches to `mask-subtract` at medium screen sizes and above.

```html
<div class="mask-add md:mask-subtract ...">  <!-- ... --></div>
```

--------------------------------

### Break Before Column Example

Source: https://tailwindcss.com/docs/break-before

Use utilities like `break-before-column` and `break-before-page` to control how a column or page break should behave before an element.

```html
<div class="columns-2">  <p>Well, let me tell you something, ...</p>  <p class="break-before-column">Sure, go ahead, laugh...</p>  <p>Maybe we can live without...</p>  <p>Look. If you think this is...</p></div>
```

--------------------------------

### Stacking multiple utilities

Source: https://tailwindcss.com/docs/font-variant-numeric

The `font-variant-numeric` utilities are composable, allowing multiple variants to be enabled by combining them. This example uses `slashed-zero` and `tabular-nums` for financial figures.

```html
<dl class="...">
  <dt class="...">Subtotal</dt>
  <dd class="text-right slashed-zero tabular-nums ...">$100.00</dd>
  <dt class="...">Tax</dt>
  <dd class="text-right slashed-zero tabular-nums ...">$14.50</dd>
  <dt class="...">Total</dt>
  <dd class="text-right slashed-zero tabular-nums ...">$114.50</dd>
</dl>
```

--------------------------------

### mask-cover utility

Source: https://tailwindcss.com/docs/mask-size

Use the `mask-cover` utility to scale the mask image until it fills the mask layer, cropping the image if needed. This example also sets a background image.

```html
<div class="mask-cover mask-[url(/img/scribble.png)] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Set Percentage-Based Height with Tailwind CSS

Source: https://tailwindcss.com/docs/height

Utilize `h-full` or `h-<fraction>` utilities for percentage-based heights. Examples include `h-1/2` and `h-2/5`.

```html
<div class="h-full ...">h-full</div><div class="h-9/10 ...">h-9/10</div><div class="h-3/4 ...">h-3/4</div><div class="h-1/2 ...">h-1/2</div><div class="h-1/3 ...">h-1/3</div>
```

--------------------------------

### Link stylesheet in TanStack Start root route

Source: https://tailwindcss.com/docs/installation/framework-guides/tanstack-start

Import the compiled CSS file into your `__root.tsx` file using the `?url` query and link it in the `head` section of your root route.

```typescript
// other imports...import appCss from '../styles.css?url'
export const Route = createRootRoute({
  head: () => ({
    meta: [
      // your meta tags and site config
    ],
    links: [{ rel: 'stylesheet', href: appCss }],
    // other head config
  }),
  component: RootComponent,
})
```

--------------------------------

### Basic Skew Utilities

Source: https://tailwindcss.com/docs/skew

Use `skew-<number>` utilities to skew an element on both axes. Examples include `skew-3`, `skew-6`, and `skew-12`.

```html
<img class="skew-3 ..." src="/img/mountains.jpg" /><img class="skew-6 ..." src="/img/mountains.jpg" /><img class="skew-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Correct Mobile-First Styling

Source: https://tailwindcss.com/docs/responsive-design

Apply styles for mobile screens using unprefixed utilities and override them for larger breakpoints. This example centers text on mobile and left-aligns it on screens 640px and wider.

```html
<!-- This will center text on mobile, and left align it on screens 640px and wider -->
<div class="text-center sm:text-left"></div>
```

--------------------------------

### Percentage-based width with fractions

Source: https://tailwindcss.com/docs/width

Use w-full or w-<fraction> utilities like w-1/2 and w-2/5 to set percentage-based widths.

```html
<div class="flex ...">  <div class="w-1/2 ...">w-1/2</div>  <div class="w-1/2 ...">w-1/2</div></div><div class="flex ...">  <div class="w-2/5 ...">w-2/5</div>  <div class="w-3/5 ...">w-3/5</div></div><div class="flex ...">  <div class="w-1/3 ...">w-1/3</div>  <div class="w-2/3 ...">w-2/3</div></div><div class="flex ...">  <div class="w-1/4 ...">w-1/4</div>  <div class="w-3/4 ...">w-3/4</div></div><div class="flex ...">  <div class="w-1/5 ...">w-1/5</div>  <div class="w-4/5 ...">w-4/5</div></div><div class="flex ...">  <div class="w-1/6 ...">w-1/6</div>  <div class="w-5/6 ...">w-5/6</div></div><div class="w-full ...">w-full</div>
```

--------------------------------

### Basic text-decoration-style HTML example

Source: https://tailwindcss.com/docs/text-decoration-style

Apply decoration style utilities to paragraphs with the underline class. Each utility (decoration-solid, decoration-double, decoration-dotted, decoration-dashed, decoration-wavy) changes the visual style of the text underline.

```html
<p class="underline decoration-solid">The quick brown fox...</p><p class="underline decoration-double">The quick brown fox...</p><p class="underline decoration-dotted">The quick brown fox...</p><p class="underline decoration-dashed">The quick brown fox...</p><p class="underline decoration-wavy">The quick brown fox...</p>
```

--------------------------------

### Scroll Padding Inline Start with Logical Properties

Source: https://tailwindcss.com/docs/scroll-padding

Use `scroll-ps-<number>` for `scroll-padding-inline-start`. This utility maps to left or right based on text direction.

```html
<div dir="ltr">  <div class="snap-x scroll-ps-6 ...">    <!-- ... -->  </div></div><div dir="rtl">  <div class="snap-x scroll-ps-6 ...">    <!-- ... -->  </div></div>
```

--------------------------------

### Apply Responsive Opacity with Tailwind CSS

Source: https://tailwindcss.com/docs/opacity

Explains how to use breakpoint variants like `md:` to apply opacity utilities responsively at different screen sizes.

```html
<button class="opacity-50 md:opacity-100 ...">  <!-- ... --></button>
```

--------------------------------

### Create new Rspack project

Source: https://tailwindcss.com/docs/installation/framework-guides/rspack/react

Initialize a new Rspack project using the official CLI tool.

```bash
npm create rspack@latest
```

--------------------------------

### Basic Backdrop Invert Examples

Source: https://tailwindcss.com/docs/backdrop-filter-invert

Use utilities like `backdrop-invert` and `backdrop-invert-65` to control the color inversion of an element's backdrop. `backdrop-invert-0` applies no inversion.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert-0 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert-65 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert ..."></div></div>
```

--------------------------------

### Avatar Display with Static Utility Classes

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Displays a list of contributor avatars using static Tailwind utility classes. This example shows potential duplication if not rendered in a loop.

```html
<div>  <div class="flex items-center space-x-2 text-base">
    <h4 class="font-semibold text-slate-900">Contributors</h4>
    <span class="bg-slate-100 px-2 py-1 text-xs font-semibold text-slate-700 ...">204</span>
  </div>
  <div class="mt-3 flex -space-x-2 overflow-hidden">
    <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" />
    <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src="https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" />
    <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src="https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2.25&w=256&h=256&q=80" alt="" />
    <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" />
    <img class="inline-block h-12 w-12 rounded-full ring-2 ring-white" src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" />
  </div>
  <div class="mt-3 text-sm font-medium">
    <a href="#" class="text-blue-500">+ 198 others</a>
  </div>
</div>
```

--------------------------------

### Create Next.js project with Create Next App

Source: https://tailwindcss.com/docs/installation/framework-guides/nextjs

Initialize a new Next.js project with TypeScript and ESLint enabled using Create Next App.

```bash
npx create-next-app@latest my-project --typescript --eslint --app
cd my-project
```

--------------------------------

### CSS accent-color property for accent-neutral-600

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-600` utility.

```css
accent-color: var(--color-neutral-600); /* oklch(43.9% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-neutral-700

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-700` utility.

```css
accent-color: var(--color-neutral-700); /* oklch(37.1% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-neutral-400

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-400` utility.

```css
accent-color: var(--color-neutral-400); /* oklch(70.8% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-zinc-600

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-600` utility.

```css
accent-color: var(--color-zinc-600); /* oklch(44.2% 0.017 285.786) */
```

--------------------------------

### CSS accent-color property for accent-neutral-900

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-900` utility.

```css
accent-color: var(--color-neutral-900); /* oklch(20.5% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-zinc-700

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-700` utility.

```css
accent-color: var(--color-zinc-700); /* oklch(37% 0.013 285.805) */
```

--------------------------------

### CSS accent-color property for accent-zinc-400

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-400` utility.

```css
accent-color: var(--color-zinc-400); /* oklch(70.5% 0.015 286.067) */
```

--------------------------------

### Responsive transition-behavior

Source: https://tailwindcss.com/docs/transition-behavior

Prefix a `transition-behavior` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This allows for different transition behaviors based on screen size.

```html
<button class="transition-discrete md:transition-normal ...">  <!-- ... --></button>
```

--------------------------------

### CSS accent-color property for accent-neutral-800

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-800` utility.

```css
accent-color: var(--color-neutral-800); /* oklch(26.9% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-neutral-950

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-950` utility.

```css
accent-color: var(--color-neutral-950); /* oklch(14.5% 0 0) */
```

--------------------------------

### Basic backdrop filter example with HTML

Source: https://tailwindcss.com/docs/backdrop-filter

Apply backdrop filter utilities like backdrop-blur-xs and backdrop-grayscale to an element's backdrop. Multiple filters can be combined on a single element.

```html
<div class="bg-[url(/img/mountains.jpg)] ...">
  <div class="backdrop-blur-xs ..."></div>
</div>
<div class="bg-[url(/img/mountains.jpg)] ...">
  <div class="backdrop-grayscale ..."></div>
</div>
<div class="bg-[url(/img/mountains.jpg)] ...">
  <div class="backdrop-blur-xs backdrop-grayscale ..."></div>
</div>
```

--------------------------------

### CSS accent-color property for accent-neutral-100

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-100` utility.

```css
accent-color: var(--color-neutral-100); /* oklch(97% 0 0) */
```

--------------------------------

### Responsive Contrast Utility

Source: https://tailwindcss.com/docs/filter-contrast

Apply contrast utilities responsively by prefixing them with a breakpoint variant, such as `md:`, to control contrast at specific screen sizes.

```html
<img class="contrast-125 md:contrast-150 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Container scale min-inline-size utilities

Source: https://tailwindcss.com/docs/min-inline-size

Use utilities like `min-inline-sm` and `min-inline-xl` to set a fixed minimum inline size based on the container scale. Examples range from `min-inline-3xs` to `min-inline-lg`.

```html
<div class="inline-40 ...">  <div class="min-inline-lg ...">min-inline-lg</div>  <div class="min-inline-md ...">min-inline-md</div>  <div class="min-inline-sm ...">min-inline-sm</div>  <div class="min-inline-xs ...">min-inline-xs</div>  <div class="min-inline-2xs ...">min-inline-2xs</div>  <div class="min-inline-3xs ...">min-inline-3xs</div></div>
```

--------------------------------

### CSS accent-color property for accent-zinc-900

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-900` utility.

```css
accent-color: var(--color-zinc-900); /* oklch(21% 0.006 285.885) */
```

--------------------------------

### CSS accent-color property for accent-zinc-800

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-800` utility.

```css
accent-color: var(--color-zinc-800); /* oklch(27.4% 0.006 286.033) */
```

--------------------------------

### justify-self-stretch Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-stretch` utility to stretch a grid item to fill the grid area on its inline axis.

```html
<div class="grid justify-items-start ...">  <!-- ... -->  <div class="justify-self-stretch ...">02</div>  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-neutral-50

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-50` utility.

```css
accent-color: var(--color-neutral-50); /* oklch(98.5% 0 0) */
```

--------------------------------

### Basic Border Widths

Source: https://tailwindcss.com/docs/border-width

Use `border` or `border-<number>` utilities to set the border width for all sides of an element. Examples include `border-2` and `border-4`.

```html
<div class="border border-indigo-600 ..."></div><div class="border-2 border-indigo-600 ..."></div><div class="border-4 border-indigo-600 ..."></div><div class="border-8 border-indigo-600 ..."></div>
```

--------------------------------

### Responsive touch-action utility in HTML

Source: https://tailwindcss.com/docs/touch-action

Use breakpoint prefixes like `md:` to apply `touch-action` utilities conditionally based on screen size, enabling different touch behaviors for various devices.

```html
<div class="touch-pan-x md:touch-auto ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-neutral-500

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-500` utility.

```css
accent-color: var(--color-neutral-500); /* oklch(55.6% 0 0) */
```

--------------------------------

### Basic min-height utilities

Source: https://tailwindcss.com/docs/min-height

Use min-h utilities to set a fixed minimum height based on the spacing scale. Examples include min-h-80, min-h-64, and min-h-24.

```html
<div class="h-20 ...">  <div class="min-h-80 ...">min-h-80</div>  <div class="min-h-64 ...">min-h-64</div>  <div class="min-h-48 ...">min-h-48</div>  <div class="min-h-40 ...">min-h-40</div>  <div class="min-h-32 ...">min-h-32</div>  <div class="min-h-24 ...">min-h-24</div></div>
```

--------------------------------

### CSS accent-color property for accent-neutral-300

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-300` utility.

```css
accent-color: var(--color-neutral-300); /* oklch(87% 0 0) */
```

--------------------------------

### Specifying grid rows with grid-rows utility

Source: https://tailwindcss.com/docs/grid-template-rows

Create a grid with equally sized rows using grid-rows-<number> utilities. The example creates a 4-row grid with gap spacing.

```html
<div class="grid grid-flow-col grid-rows-4 gap-4">
  <div>01</div>
  <!-- ... -->
  <div>09</div>
</div>
```

--------------------------------

### CSS accent-color property for accent-neutral-200

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-neutral-200` utility.

```css
accent-color: var(--color-neutral-200); /* oklch(92.2% 0 0) */
```

--------------------------------

### CSS accent-color property for accent-zinc-100

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-100` utility.

```css
accent-color: var(--color-zinc-100); /* oklch(96.7% 0.001 286.375) */
```

--------------------------------

### Applying responsive border styles in HTML

Source: https://tailwindcss.com/docs/border-style

Demonstrates how to apply different border styles based on screen size using responsive variants like `md:border-dotted`.

```html
<div class="border-solid md:border-dotted ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-zinc-950

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-950` utility.

```css
accent-color: var(--color-zinc-950); /* oklch(14.1% 0.005 285.823) */
```

--------------------------------

### CSS accent-color property for accent-zinc-500

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-500` utility.

```css
accent-color: var(--color-zinc-500); /* oklch(55.2% 0.016 285.938) */
```

--------------------------------

### CSS accent-color property for accent-zinc-300

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-300` utility.

```css
accent-color: var(--color-zinc-300); /* oklch(87.1% 0.006 286.286) */
```

--------------------------------

### Apply responsive padding with breakpoint variant

Source: https://tailwindcss.com/docs/padding

Use the md: breakpoint prefix to apply padding utilities only at medium screen sizes and above. This example applies py-4 on smaller screens and py-8 on medium screens and larger.

```html
<div class="py-4 md:py-8 ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-zinc-50

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-50` utility.

```css
accent-color: var(--color-zinc-50); /* oklch(98.5% 0 0) */
```

--------------------------------

### Responsive Hyphenation

Source: https://tailwindcss.com/docs/hyphens

Apply hyphenation utilities responsively by prefixing them with a breakpoint variant, such as `md:`. This example applies `hyphens-none` by default and `hyphens-auto` at medium screen sizes and above.

```html
<p class="hyphens-none md:hyphens-auto ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Applying accent-color utility

Source: https://tailwindcss.com/docs/accent-color

Use the `accent-{color}` utility to set the accent color for form elements. This example shows how to apply `accent-gray-900`.

```html
<input class="accent-gray-900" type="checkbox">
<input class="accent-gray-900" type="radio" checked>
<input class="accent-gray-900" type="range">
```

--------------------------------

### Setting Inline Border Color to Green Shades

Source: https://tailwindcss.com/docs/border-color

This example demonstrates applying inline border colors using different shades of green, from `green-50` to `green-950`.

```html
<!-- Set border-inline-color to green-50 -->
<div class="border-x-green-50"></div>

<!-- Set border-inline-color to green-100 -->
<div class="border-x-green-100"></div>

<!-- Set border-inline-color to green-200 -->
<div class="border-x-green-200"></div>

<!-- Set border-inline-color to green-300 -->
<div class="border-x-green-300"></div>

<!-- Set border-inline-color to green-400 -->
<div class="border-x-green-400"></div>

<!-- Set border-inline-color to green-500 -->
<div class="border-x-green-500"></div>

<!-- Set border-inline-color to green-600 -->
<div class="border-x-green-600"></div>

<!-- Set border-inline-color to green-700 -->
<div class="border-x-green-700"></div>

<!-- Set border-inline-color to green-800 -->
<div class="border-x-green-800"></div>

<!-- Set border-inline-color to green-900 -->
<div class="border-x-green-900"></div>

<!-- Set border-inline-color to green-950 -->
<div class="border-x-green-950"></div>
```

--------------------------------

### CSS accent-color property for accent-zinc-200

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-zinc-200` utility.

```css
accent-color: var(--color-zinc-200); /* oklch(92% 0.004 286.32) */
```

--------------------------------

### Percentage-based min-inline-size utilities

Source: https://tailwindcss.com/docs/min-inline-size

Use `min-inline-full` or `min-inline-<fraction>` utilities like `min-inline-1/2` and `min-inline-2/5` to set a percentage-based minimum inline size.

```html
<div class="flex ...">  <div class="min-inline-3/4 ...">min-inline-3/4</div>  <div class="inline-full ...">inline-full</div></div>
```

--------------------------------

### Scroll Padding Block Start with Logical Properties

Source: https://tailwindcss.com/docs/scroll-padding

Use `scroll-pbs-<number>` for `scroll-padding-block-start`. This utility maps to top or bottom based on writing mode.

```html
<div class="snap-y scroll-pbs-6 ...">  <!-- ... --></div>
```

--------------------------------

### CSS accent-color property for accent-gray-950

Source: https://tailwindcss.com/docs/accent-color

This example shows the CSS generated for the `accent-gray-950` utility.

```css
accent-color: var(--color-gray-950); /* oklch(13% 0.028 261.692) */
```

--------------------------------

### Configure Prefix and Theme Variables

Source: https://tailwindcss.com/docs/upgrade-guide

Import Tailwind with a prefix and define theme variables. Generated CSS variables will include the prefix.

```css
@import "tailwindcss" prefix(tw);@theme {
  --font-display: "Satoshi", "sans-serif";
  --breakpoint-3xl: 120rem;
  --color-avocado-100: oklch(0.99 0 0);
  --color-avocado-200: oklch(0.98 0.04 113.22);
  --color-avocado-300: oklch(0.94 0.11 115.03);
  /* ... */
}
```

```css
:root {
  --tw-font-display: "Satoshi", "sans-serif";
  --tw-breakpoint-3xl: 120rem;
  --tw-color-avocado-100: oklch(0.99 0 0);
  --tw-color-avocado-200: oklch(0.98 0.04 113.22);
  --tw-color-avocado-300: oklch(0.94 0.11 115.03);
  /* ... */
}
```

--------------------------------

### justify-self-center Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-center` utility to align a grid item along the center of its inline axis.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-center ...">02</div>  <!-- ... --></div>
```

--------------------------------

### justify-items-center Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-center` utility to justify grid items against the end of their inline axis.

```html
<div class="grid grid-flow-col justify-items-center ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Basic Backdrop Grayscale Example

Source: https://tailwindcss.com/docs/backdrop-filter-grayscale

Use utilities like `backdrop-grayscale-50` and `backdrop-grayscale` to control the grayscale effect applied to an element's backdrop. `backdrop-grayscale-0` removes the effect.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-grayscale-0 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-grayscale-50 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-grayscale-200 ..."></div></div>
```

--------------------------------

### Static Positioning Example

Source: https://tailwindcss.com/docs/position

Use the `static` utility to position an element according to the normal flow of the document. Offsets are ignored, and the element does not act as a position reference for absolutely positioned children.

```html
<div class="static ...">
  <p>Static parent</p>
  <div class="absolute bottom-0 left-0 ...">
    <p>Absolute child</p>
  </div>
</div>
```

--------------------------------

### justify-items-end Example

Source: https://tailwindcss.com/docs/justify-items

Use the `justify-items-end` utility to justify grid items against the end of their inline axis.

```html
<div class="grid grid-flow-col justify-items-end ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Using Container Query Units

Source: https://tailwindcss.com/docs/responsive-design

Reference container sizes using container query length units like `cqw` as arbitrary values in utility classes for dynamic sizing.

```html
<div class="@container">  <div class="w-[50cqw]">    <!-- ... -->  </div></div>
```

--------------------------------

### Basic Element Scaling

Source: https://tailwindcss.com/docs/scale

Use `scale-<number>` utilities to scale an element by a percentage of its original size. Examples include `scale-75`, `scale-100`, and `scale-125`.

```html
<img class="scale-75 ..." src="/img/mountains.jpg" /><img class="scale-100 ..." src="/img/mountains.jpg" /><img class="scale-125 ..." src="/img/mountains.jpg" />
```

--------------------------------

### justify-self-end Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-end` utility to align a grid item to the end of its inline axis.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-end ...">02</div>  <!-- ... --></div>
```

--------------------------------

### Basic max-block-size Utilities

Source: https://tailwindcss.com/docs/max-block-size

Use `max-block-<number>` utilities to set a fixed maximum block size based on the spacing scale. Examples include `max-block-80`, `max-block-64`, `max-block-48`, `max-block-40`, and `max-block-32`.

```html
<div class="block-96 ...">  <div class="block-full max-block-80 ...">max-block-80</div>  <div class="block-full max-block-64 ...">max-block-64</div>  <div class="block-full max-block-48 ...">max-block-48</div>  <div class="block-full max-block-40 ...">max-block-40</div>  <div class="block-full max-block-32 ...">max-block-32</div></div>
```

--------------------------------

### justify-self-auto Example

Source: https://tailwindcss.com/docs/justify-self

Use the `justify-self-auto` utility to align an item based on the value of the grid's `justify-items` property.

```html
<div class="grid justify-items-stretch ...">  <!-- ... -->  <div class="justify-self-auto ...">02</div>  <!-- ... --></div>
```

--------------------------------

### Responsive Height Design with Tailwind CSS

Source: https://tailwindcss.com/docs/height

Prefix height utilities with breakpoint variants like `md:` to apply them only at specific screen sizes and above. Example: `md:h-full`.

```html
<div class="h-1/2 md:h-full ...">  <!-- ... --></div>
```

--------------------------------

### Basic Backdrop Blur Examples

Source: https://tailwindcss.com/docs/backdrop-filter-blur

Apply different levels of backdrop blur to elements using predefined utilities like backdrop-blur-sm and backdrop-blur-md. These utilities are applied to elements with a background image and a semi-transparent white overlay.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-blur-none ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-blur-sm ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-blur-md ..."></div></div>
```

--------------------------------

### Responsive scroll-snap-stop

Source: https://tailwindcss.com/docs/scroll-snap-stop

Apply `scroll-snap-stop` utilities conditionally based on screen size using responsive prefixes like `md:`. This example applies `snap-always` by default and switches to `snap-normal` at medium screen sizes and above.

```html
<div class="snap-always md:snap-normal ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive min-inline-size design

Source: https://tailwindcss.com/docs/min-inline-size

Prefix a `min-inline-size` utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above. For example, `md:min-inline-0`.

```html
<div class="inline-24 min-inline-full md:min-inline-0 ...">  <!-- ... --></div>
```

--------------------------------

### Fixed Positioning Example

Source: https://tailwindcss.com/docs/position

Use the `fixed` utility to position an element relative to the browser window. Offsets are calculated relative to the viewport.

```html
<div class="relative">
  <div class="fixed top-0 right-0 left-0">Contacts</div>
  <div>
    <div>
      <img src="/img/andrew.jpg" />
      <strong>Andrew Alfred</strong>
    </div>
    <div>
      <img src="/img/debra.jpg" />
      <strong>Debra Houston</strong>
    </div>
    <!-- ... -->
  </div>
</div>
```

--------------------------------

### Applying Responsive Perspective Origin Utilities in HTML

Source: https://tailwindcss.com/docs/perspective-origin

Control perspective origin responsively by prefixing utilities with breakpoint variants like `md:` to apply different styles at medium screen sizes and above.

```html
<div class="perspective-origin-center md:perspective-origin-bottom-left ...">  <!-- ... --></div>
```

--------------------------------

### Set Column Gap

Source: https://tailwindcss.com/docs/columns

Use `gap-<width>` utilities to control the space between columns. This example sets a gap of `8` units.

```html
<div class="columns-3 gap-8 ...">
  <img class="aspect-3/2 ..." src="/img/mountains-1.jpg" />
  <img class="aspect-square ..." src="/img/mountains-2.jpg" />
  <img class="aspect-square ..." src="/img/mountains-3.jpg" />
  <!-- ... -->
</div>
```

--------------------------------

### Responsive mix-blend-mode Design

Source: https://tailwindcss.com/docs/mix-blend-mode

Prefix a `mix-blend-mode` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example applies `mix-blend-multiply` by default and `mix-blend-overlay` from the medium breakpoint onwards.

```html
<div class="mix-blend-multiply md:mix-blend-overlay ...">  <!-- ... --></div>
```

--------------------------------

### Basic color-scheme HTML example

Source: https://tailwindcss.com/docs/color-scheme

Apply color-scheme utilities to control how elements render across different system color schemes. Use scheme-light, scheme-dark, or scheme-light-dark to set the color scheme for date inputs or other form elements.

```html
<div class="scheme-light ...">  <input type="date" /></div><div class="scheme-dark ...">  <input type="date" /></div><div class="scheme-light-dark ...">  <input type="date" /></div>
```

--------------------------------

### Video Aspect Ratio

Source: https://tailwindcss.com/docs/aspect-ratio

Use the `aspect-video` utility to give a video element a 16 / 9 aspect ratio. Resize the example to see the expected behavior.

```html
<iframe class="aspect-video ..." src="https://www.youtube.com/embed/dQw4w9WgXcQ"></iframe>
```

--------------------------------

### Set Border-block Color with Taupe Palette

Source: https://tailwindcss.com/docs/border-color

This example demonstrates setting the `border-block-color` using a shade from the taupe palette.

```css
border-block-color: var(--color-taupe-950); /* oklch(14.7% 0.004 49.3) */
```

--------------------------------

### Configure assets.deploy alias in mix.exs

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Set up the deployment alias to build and minify CSS and JavaScript assets during deployment.

```elixir
defp aliases do  [    # …    "assets.deploy": [      "tailwind myproject --minify",      "esbuild myproject --minify",      "phx.digest"    ]  ]end
```

--------------------------------

### Percentage-based min-width

Source: https://tailwindcss.com/docs/min-width

Use min-w-full or fractional utilities like min-w-1/2 and min-w-2/5 to set percentage-based minimum widths.

```html
<div class="flex ...">
  <div class="min-w-3/4 ...">min-w-3/4</div>
  <div class="w-full ...">w-full</div>
</div>
```

--------------------------------

### Using custom grid-auto-rows value

Source: https://tailwindcss.com/docs/grid-auto-rows

Shows how to set grid row size using the `auto-rows-[<value>]` syntax for a custom `minmax` value.

```html
<div class="auto-rows-[minmax(0,2fr)] ...">  <!-- ... --></div>
```

--------------------------------

### Responsive brightness with breakpoint variants

Source: https://tailwindcss.com/docs/filter-brightness

Combine brightness utilities with breakpoint prefixes like md: to apply different brightness levels at different screen sizes.

```html
<img class="brightness-110 md:brightness-150 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Negative Margin Utilities

Source: https://tailwindcss.com/docs/margin

Prefix a margin utility with a dash to use a negative margin value. This example shows a negative top margin.

```html
<div class="h-16 w-36 bg-sky-400 opacity-20 ..."></div><div class="-mt-8 bg-sky-300 ...">-mt-8</div>
```

--------------------------------

### Starting and ending grid lines with col-start and col-end

Source: https://tailwindcss.com/docs/grid-column

Utilize `col-start-<number>` and `col-end-<number>` utilities to position elements at specific grid lines. These can be combined with `col-span-<number>` for precise control.

```html
<div class="grid grid-cols-6 gap-4">
  <div class="col-span-4 col-start-2 ...">01</div>
  <div class="col-start-1 col-end-3 ...">02</div>
  <div class="col-span-2 col-end-7 ...">03</div>
  <div class="col-start-1 col-end-7 ...">04</div>
</div>
```

--------------------------------

### Conic Gradient From Position

Source: https://tailwindcss.com/docs/mask-image

Set the starting position for the conic gradient mask. Can be defined by degrees, percentages, colors, or custom properties.

```css
mask-image: conic-gradient(from var(--tw-mask-conic-position), black calc(var(--spacing) * <number>), transparent var(--tw-mask-conic-to));
```

```css
mask-image: conic-gradient(from var(--tw-mask-conic-position), black <percentage>, transparent var(--tw-mask-conic-to));
```

```css
mask-image: conic-gradient(from var(--tw-mask-conic-position), <color> var(--tw-mask-conic-from), transparent var(--tw-mask-conic-to));
```

```css
mask-image: conic-gradient(from var(--tw-mask-conic-position), black var(<custom-property>), transparent var(--tw-mask-conic-to));
```

```css
mask-image: conic-gradient(from var(--tw-mask-conic-position), black <value>, transparent var(--tw-mask-conic-to));
```

--------------------------------

### Set border-block-color to lime shades

Source: https://tailwindcss.com/docs/border-color

This example demonstrates setting the `border-block-color` using CSS variables for different lime shades, from 50 to 950.

```css
.border-y-lime-50 { border-block-color: var(--color-lime-50); /* oklch(98.6% 0.031 120.757) */ }
```

```css
.border-y-lime-100 { border-block-color: var(--color-lime-100); /* oklch(96.7% 0.067 122.328) */ }
```

```css
.border-y-lime-200 { border-block-color: var(--color-lime-200); /* oklch(93.8% 0.127 124.321) */ }
```

```css
.border-y-lime-300 { border-block-color: var(--color-lime-300); /* oklch(89.7% 0.196 126.665) */ }
```

```css
.border-y-lime-400 { border-block-color: var(--color-lime-400); /* oklch(84.1% 0.238 128.85) */ }
```

```css
.border-y-lime-500 { border-block-color: var(--color-lime-500); /* oklch(76.8% 0.233 130.85) */ }
```

```css
.border-y-lime-600 { border-block-color: var(--color-lime-600); /* oklch(64.8% 0.2 131.684) */ }
```

```css
.border-y-lime-700 { border-block-color: var(--color-lime-700); /* oklch(53.2% 0.157 131.589) */ }
```

```css
.border-y-lime-800 { border-block-color: var(--color-lime-800); /* oklch(45.3% 0.124 130.933) */ }
```

```css
.border-y-lime-900 { border-block-color: var(--color-lime-900); /* oklch(40.5% 0.101 131.063) */ }
```

```css
.border-y-lime-950 { border-block-color: var(--color-lime-950); /* oklch(27.4% 0.072 132.109) */ }
```

--------------------------------

### Set minimum block size with percentages

Source: https://tailwindcss.com/docs/min-block-size

Use `min-block-full` or `min-block-<fraction>` utilities to give an element a percentage-based minimum block size.

```html
<div class="min-block-full ...">min-block-full</div><div class="min-block-9/10 ...">min-block-9/10</div><div class="min-block-3/4 ...">min-block-3/4</div><div class="min-block-1/2 ...">min-block-1/2</div><div class="min-block-1/3 ...">min-block-1/3</div>
```

--------------------------------

### Animating with CSS Variables in React

Source: https://tailwindcss.com/docs/upgrade-guide

This React example demonstrates using CSS variables directly for animation properties, leveraging the simplicity and performance benefits of v4.

```jsx
<motion.div animate={{ backgroundColor: "var(--color-blue-500)" }} />
```

--------------------------------

### Basic Positioning Utilities

Source: https://tailwindcss.com/docs/top-right-bottom-left

Use these utilities to set the horizontal or vertical position of a positioned element. Examples include pinning to corners, spanning edges, and filling the parent.

```html
<!-- Pin to top left corner --><div class="relative size-32 ...">  <div class="absolute top-0 left-0 size-16 ...">01</div></div><!-- Span top edge --><div class="relative size-32 ...">  <div class="absolute inset-x-0 top-0 h-16 ...">02</div></div><!-- Pin to top right corner --><div class="relative size-32 ...">  <div class="absolute top-0 right-0 size-16 ...">03</div></div><!-- Span left edge --><div class="relative size-32 ...">  <div class="absolute inset-y-0 left-0 w-16 ...">04</div></div><!-- Fill entire parent --><div class="relative size-32 ...">  <div class="absolute inset-0 ...">05</div></div><!-- Span right edge --><div class="relative size-32 ...">  <div class="absolute inset-y-0 right-0 w-16 ...">06</div></div><!-- Pin to bottom left corner --><div class="relative size-32 ...">  <div class="absolute bottom-0 left-0 size-16 ...">07</div></div><!-- Span bottom edge --><div class="relative size-32 ...">  <div class="absolute inset-x-0 bottom-0 h-16 ...">08</div></div><!-- Pin to bottom right corner --><div class="relative size-32 ...">  <div class="absolute right-0 bottom-0 size-16 ...">09</div></div>
```

--------------------------------

### Basic transition-delay utilities

Source: https://tailwindcss.com/docs/transition-delay

Use utilities like `delay-150` and `delay-700` to set the transition delay of an element in milliseconds. These classes are applied to buttons to demonstrate the effect.

```html
<button class="transition delay-150 duration-300 ease-in-out ...">Button A</button><button class="transition delay-300 duration-300 ease-in-out ...">Button B</button><button class="transition delay-700 duration-300 ease-in-out ...">Button C</button>
```

--------------------------------

### Conditional transition-delay with reduced motion

Source: https://tailwindcss.com/docs/transition-delay

Conditionally apply animations and transitions for users who prefer reduced motion using `motion-safe` and `motion-reduce` variants. This example sets the delay to 0 when reduced motion is detected.

```html
<button type="button" class="delay-300 motion-reduce:delay-0 ...">  <!-- ... --></button>
```

--------------------------------

### Transition with reduced motion support

Source: https://tailwindcss.com/docs/transition-property

Conditionally apply animations and transitions using the `motion-safe` and `motion-reduce` variants for users who prefer reduced motion. This example disables transitions on hover for users with reduced motion preferences.

```html
<button class="transform transition hover:-translate-y-1 motion-reduce:transition-none motion-reduce:hover:transform-none ...">  <!-- ... --></button>
```

--------------------------------

### Start-start logical border radius with preset sizes

Source: https://tailwindcss.com/docs/border-radius

Apply border-radius to the start-start corner using logical properties. Supports sizes from xs (0.125rem) to 4xl (2rem).

```css
border-start-start-radius: var(--radius-xs); /* 0.125rem (2px) */
```

```css
border-start-start-radius: var(--radius-sm); /* 0.25rem (4px) */
```

```css
border-start-start-radius: var(--radius-md); /* 0.375rem (6px) */
```

```css
border-start-start-radius: var(--radius-lg); /* 0.5rem (8px) */
```

```css
border-start-start-radius: var(--radius-xl); /* 0.75rem (12px) */
```

```css
border-start-start-radius: var(--radius-2xl); /* 1rem (16px) */
```

```css
border-start-start-radius: var(--radius-3xl); /* 1.5rem (24px) */
```

```css
border-start-start-radius: var(--radius-4xl); /* 2rem (32px) */
```

--------------------------------

### Specify Grid Columns

Source: https://tailwindcss.com/docs/grid-template-columns

Use grid-cols-n utilities to create grids with n equally sized columns. For example, grid-cols-4 creates a 4-column grid.

```html
<div class="grid grid-cols-4 gap-4">  <div>01</div>  <!-- ... -->  <div>09</div></div>
```

--------------------------------

### Custom transition-delay value

Source: https://tailwindcss.com/docs/transition-delay

Use the `delay-[<value>]` syntax to set the transition delay with a completely custom value. This example demonstrates setting a delay of '1s,250ms'.

```html
<button class="delay-[1s,250ms] ...">  <!-- ... --></button>
```

--------------------------------

### Using a Custom Utility in HTML

Source: https://tailwindcss.com/docs/adding-custom-styles

Demonstrates how to use a custom utility class defined with @utility in HTML markup.

```html
<div class="content-auto">
  <!-- ... -->
</div>
```

--------------------------------

### Basic Rotation with Tailwind CSS

Source: https://tailwindcss.com/docs/rotate

Use `rotate-<number>` utilities like `rotate-45` and `rotate-90` to rotate an element by degrees. Examples include 45, 90, and 210 degrees.

```html
<img class="rotate-45 ..." src="/img/mountains.jpg" /><img class="rotate-90 ..." src="/img/mountains.jpg" /><img class="rotate-210 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Use Theme Variables in Arbitrary Values with calc()

Source: https://tailwindcss.com/docs/theme

Combine theme variables with calc() in arbitrary values for dynamic calculations. In this example, subtracting 1px from a radius variable creates a concentric border effect.

```HTML
<div class="relative rounded-xl">
  <div class="absolute inset-px rounded-[calc(var(--radius-xl)-1px)]">
    <!-- ... -->
  </div>
  <!-- ... -->
</div>
```

--------------------------------

### Apply Responsive table-layout Utilities

Source: https://tailwindcss.com/docs/table-layout

Use breakpoint prefixes like `md:` to apply `table-layout` utilities only at specific screen sizes and above. This allows for adaptive table designs.

```html
<div class="table-auto md:table-fixed ...">  <!-- ... --></div>
```

--------------------------------

### Flex Items with Percentage Basis

Source: https://tailwindcss.com/docs/flex-basis

Sets the initial size of flex items using percentage-based fractions. Utilities like `basis-1/2` and `basis-2/3` are used.

```html
<div class="flex flex-row">  <div class="basis-1/3">01</div>  <div class="basis-2/3">02</div></div>
```

--------------------------------

### Negative Skew Utilities

Source: https://tailwindcss.com/docs/skew

Use `-skew-<number>` utilities to skew an element on both axes with negative values. Examples include `-skew-3`, `-skew-6`, and `-skew-12`.

```html
<img class="-skew-3 ..." src="/img/mountains.jpg" /><img class="-skew-6 ..." src="/img/mountains.jpg" /><img class="-skew-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Skew X-axis Utilities

Source: https://tailwindcss.com/docs/skew

Use `skew-x-<number>` and `-skew-x-<number>` utilities to skew an element specifically on the x-axis. Examples include `skew-x-6` and `-skew-x-12`.

```html
<img class="-skew-x-12 ..." src="/img/mountains.jpg" /><img class="skew-x-6 ..." src="/img/mountains.jpg" /><img class="skew-x-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Use a Custom Variant in HTML

Source: https://tailwindcss.com/docs/adding-custom-styles

Apply the newly defined custom variant to HTML elements by prefixing utility classes with the variant name. This example shows how to use the `theme-midnight` variant.

```html
<html data-theme="midnight">
  <button class="theme-midnight:bg-black ..."></button>
</html>
```

--------------------------------

### Apply Tailwind CSS classes in a component

Source: https://tailwindcss.com/docs/installation/framework-guides/tanstack-start

Demonstrates how to use Tailwind's utility classes directly in a React component within a TanStack Start project.

```typescript
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <h1 class="text-3xl font-bold underline">
      Hello World!
    </h1>
  )
}
```

--------------------------------

### Start-start logical border radius with arbitrary value

Source: https://tailwindcss.com/docs/border-radius

Apply start-start border-radius using an arbitrary CSS value.

```css
border-start-start-radius: <value>;
```

--------------------------------

### Sticky Positioning Example

Source: https://tailwindcss.com/docs/position

Use the `sticky` utility to position an element as `relative` until it crosses a threshold, then as `fixed` until its parent is off-screen. Offsets are relative to the element's normal position.

```html
<div>
  <div>
    <div class="sticky top-0 ...">A</div>
    <div>
      <div>
        <img src="/img/andrew.jpg" />
        <strong>Andrew Alfred</strong>
      </div>
      <div>
        <img src="/img/aisha.jpg" />
        <strong>Aisha Houston</strong>
      </div>
      <!-- ... -->
    </div>
  </div>
  <div>
    <div class="sticky top-0">B</div>
    <div>
      <div>
        <img src="/img/bob.jpg" />
        <strong>Bob Alfred</strong>
      </div>
      <!-- ... -->
    </div>
  </div>
  <!-- ... -->
</div>
```

--------------------------------

### Apply inline start/end padding with ps/pe-<number>

Source: https://tailwindcss.com/docs/padding

Use `ps-<number>` or `pe-<number>` utilities to set the `padding-inline-start` and `padding-inline-end` logical properties, which map to either the left or right side based on the text direction.

```html
<div>  <div dir="ltr">    <div class="ps-8 ...">ps-8</div>    <div class="pe-8 ...">pe-8</div>  </div>  <div dir="rtl">    <div class="ps-8 ...">ps-8</div>    <div class="pe-8 ...">pe-8</div>  </div></div>
```

--------------------------------

### Block Start Border Colors (Arbitrary Values)

Source: https://tailwindcss.com/docs/border-color

Set the `border-block-start-color` using arbitrary values. This allows for direct use of CSS custom properties or any valid color value.

```html
<!-- Using a custom property -->
<div class="border-bs-[--my-color-variable] ..."></div>

<!-- Using an arbitrary value -->
<div class="border-bs-[oklch(45%_0.1_255)] ..."></div>
```

--------------------------------

### Skew Y-axis Utilities

Source: https://tailwindcss.com/docs/skew

Use `skew-y-<number>` and `-skew-y-<number>` utilities to skew an element specifically on the y-axis. Examples include `skew-y-6` and `-skew-y-12`.

```html
<img class="-skew-y-12 ..." src="/img/mountains.jpg" /><img class="skew-y-6 ..." src="/img/mountains.jpg" /><img class="skew-y-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Wrap Flex Items in Reverse Direction

Source: https://tailwindcss.com/docs/flex-wrap

Use `flex-wrap-reverse` to wrap flex items in the reverse direction, starting new lines from bottom to top.

```html
<div class="flex flex-wrap-reverse">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Update Shadow, Radius, and Blur Utilities

Source: https://tailwindcss.com/docs/upgrade-guide

Replace v3 utilities with their v4 equivalents to adapt to the updated scales for shadow, radius, and blur.

```html
<input class="shadow-sm" /><input class="shadow-xs" /><input class="shadow" /><input class="shadow-sm" />
```

--------------------------------

### Run Tailwind CSS Upgrade Tool

Source: https://tailwindcss.com/docs/upgrade-guide

Use this command to automate the migration of your project from Tailwind CSS v3 to v4. Ensure Node.js 20 or higher is installed. It's recommended to run this in a new branch.

```bash
$ npx @tailwindcss/upgrade

```

--------------------------------

### Set Border-inline-start Color with Orange Palette

Source: https://tailwindcss.com/docs/border-color

Demonstrates setting the inline-start border color using various shades from the orange color palette.

```css
border-inline-start-color: var(--color-orange-800); /* oklch(47% 0.157 37.304) */
```

--------------------------------

### Conic Gradient Mask

Source: https://tailwindcss.com/docs/mask-image

Use conic-gradient for mask images. Customize the angle, start point, and end point.

```css
mask-image: conic-gradient(from <number>deg, black var(--tw-mask-conic-from), transparent var(--tw-mask-conic-to));
```

```css
mask-image: conic-gradient(from calc(<number>deg * -1), black var(--tw-mask-conic-from), transparent var(--tw-mask-conic-to));
```

--------------------------------

### Basic grid-auto-columns Usage

Source: https://tailwindcss.com/docs/grid-auto-columns

Demonstrates controlling implicit grid column size with `auto-cols-min` and `auto-cols-max` utilities. Use this when you need to define the minimum or maximum size of columns that are not explicitly defined.

```html
<div class="grid auto-cols-max grid-flow-col">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Customizing Backdrop Blur Theme

Source: https://tailwindcss.com/docs/backdrop-filter-blur

Customize the backdrop blur utilities by defining `--blur-*` theme variables. This example shows how to set `--blur-2xs` to `2px`, making the `backdrop-blur-2xs` utility available.

```css
@theme {  --blur-2xs: 2px; }
```

--------------------------------

### Setting All Borders Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-{color}-{shade}` utility to set the color of all four borders simultaneously. For example, `border-indigo-500` sets all borders to a medium indigo.

```html
<div class="border-indigo-500">...</div>
```

--------------------------------

### Basic transition duration with multiple buttons

Source: https://tailwindcss.com/docs/transition-duration

Apply transition duration utilities to buttons with different millisecond values. Combine with transition and ease-in-out classes for smooth animations.

```html
<button class="transition duration-150 ease-in-out ...">Button A</button><button class="transition duration-300 ease-in-out ...">Button B</button><button class="transition duration-700 ease-in-out ...">Button C</button>
```

--------------------------------

### Set Caret Color with Tailwind CSS (Taupe)

Source: https://tailwindcss.com/docs/caret-color

This example demonstrates using the `caret-{color}` utility with predefined taupe colors to style the text cursor.

```html
caret-taupe-50
caret-color: var(--color-taupe-50); /* oklch(98.6% 0.002 67.8) */
caret-taupe-100
caret-color: var(--color-taupe-100); /* oklch(96% 0.002 17.2) */
caret-taupe-200
caret-color: var(--color-taupe-200); /* oklch(92.2% 0.005 34.3) */
caret-taupe-300
caret-color: var(--color-taupe-300); /* oklch(86.8% 0.007 39.5) */
caret-taupe-400
caret-color: var(--color-taupe-400); /* oklch(71.4% 0.014 41.2) */
caret-taupe-500
caret-color: var(--color-taupe-500); /* oklch(54.7% 0.021 43.1) */
caret-taupe-600
caret-color: var(--color-taupe-600); /* oklch(43.8% 0.017 39.3) */
caret-taupe-700
caret-color: var(--color-taupe-700); /* oklch(36.7% 0.016 35.7) */
caret-taupe-800
caret-color: var(--color-taupe-800); /* oklch(26.8% 0.011 36.5) */
caret-taupe-900
caret-color: var(--color-taupe-900); /* oklch(21.4% 0.009 43.1) */
caret-taupe-950
caret-color: var(--color-taupe-950); /* oklch(14.7% 0.004 49.3) */
```

--------------------------------

### Applying Responsive Display Utilities

Source: https://tailwindcss.com/docs/display

Prefix display utilities with breakpoint variants like `md:` to apply them only at specific screen sizes and above. This enables responsive design by conditionally altering element display.

```html
<div class="flex md:inline-flex ...">
  <!-- ... -->
</div>
```

--------------------------------

### 3D Rotation in Tailwind CSS

Source: https://tailwindcss.com/docs/rotate

Combine `rotate-x-<number>`, `rotate-y-<number>`, and `rotate-z-<number>` utilities for 3D rotations. Examples include `rotate-x-50`, `-rotate-y-30`, and `rotate-z-45`.

```html
<img class="rotate-x-50 rotate-z-45 ..." src="/img/mountains.jpg" /><img class="rotate-x-15 -rotate-y-30 ..." src="/img/mountains.jpg" /><img class="rotate-y-25 rotate-z-30 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Apply Prefix to Utilities

Source: https://tailwindcss.com/docs/upgrade-guide

Prefixes are now applied at the beginning of the class name, similar to variants. Configure theme variables as if no prefix is used.

```html
<div class="tw:flex tw:bg-red-500 tw:hover:bg-red-600">  <!-- ... --></div>
```

--------------------------------

### HTML Preformatted Text with Tab Sizes

Source: https://tailwindcss.com/docs/tab-size

Demonstrates how to apply `tab-2` and `tab-8` classes to `<pre>` elements to visually represent different tab sizes.

```html
<pre class="tab-2 ...">function indent() {
	return 'tabbed'
}</pre><pre class="tab-8 ...">function indent() {
	return 'tabbed'
}</pre>
```

--------------------------------

### Apply responsive order utilities in Tailwind CSS

Source: https://tailwindcss.com/docs/order

Use responsive prefixes like `md:` with order utilities to control item order based on different screen sizes.

```html
<div class="order-first md:order-last ...">  <!-- ... --></div>
```

--------------------------------

### Responsive flex utilities

Source: https://tailwindcss.com/docs/flex

Prefix a `flex` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="flex-none md:flex-1 ...">  <!-- ... --></div>
```

--------------------------------

### Apply box-sizing utilities responsively

Source: https://tailwindcss.com/docs/box-sizing

Prefix a `box-sizing` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. Learn more about using variants in the variants documentation.

```html
<div class="box-content md:box-border ...">  <!-- ... --></div>
```

--------------------------------

### Incorrect Mobile Styling with `sm:` Prefix

Source: https://tailwindcss.com/docs/responsive-design

Avoid using `sm:` prefixes for styling mobile screens. This example incorrectly applies text centering only on screens 640px and wider, not on smaller mobile devices.

```html
<!-- This will only center text on screens 640px and wider, not on small screens -->
<div class="sm:text-center"></div>
```

--------------------------------

### Set border-block-color to amber shades

Source: https://tailwindcss.com/docs/border-color

Demonstrates the use of CSS variables for setting `border-block-color` with various amber shades. Each example corresponds to a specific shade from 50 to 950.

```css
.border-y-amber-50 { border-block-color: var(--color-amber-50); /* oklch(98.7% 0.022 95.277) */ }
```

```css
.border-y-amber-100 { border-block-color: var(--color-amber-100); /* oklch(96.2% 0.059 95.617) */ }
```

```css
.border-y-amber-200 { border-block-color: var(--color-amber-200); /* oklch(92.4% 0.12 95.746) */ }
```

```css
.border-y-amber-300 { border-block-color: var(--color-amber-300); /* oklch(87.9% 0.169 91.605) */ }
```

```css
.border-y-amber-400 { border-block-color: var(--color-amber-400); /* oklch(82.8% 0.189 84.429) */ }
```

```css
.border-y-amber-500 { border-block-color: var(--color-amber-500); /* oklch(76.9% 0.188 70.08) */ }
```

```css
.border-y-amber-600 { border-block-color: var(--color-amber-600); /* oklch(66.6% 0.179 58.318) */ }
```

```css
.border-y-amber-700 { border-block-color: var(--color-amber-700); /* oklch(55.5% 0.163 48.998) */ }
```

```css
.border-y-amber-800 { border-block-color: var(--color-amber-800); /* oklch(47.3% 0.137 46.201) */ }
```

```css
.border-y-amber-900 { border-block-color: var(--color-amber-900); /* oklch(41.4% 0.112 45.904) */ }
```

```css
.border-y-amber-950 { border-block-color: var(--color-amber-950); /* oklch(27.9% 0.077 45.635) */ }
```

--------------------------------

### Responsive Grid Columns

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply different column counts for grids based on screen size using responsive variants like `md` and `lg`. This example shows a 3-column grid on mobile, 4 on medium, and 6 on large screens.

```html
<div class="grid grid-cols-3 md:grid-cols-4 lg:grid-cols-6">  <!-- ... --></div>
```

--------------------------------

### Absolute Positioning Example

Source: https://tailwindcss.com/docs/position

Use the `absolute` utility to position an element outside the normal document flow. Offsets are relative to the nearest parent with a non-`static` position.

```html
<div class="static ...">
  <!-- Static parent -->
  <div class="static ..."><p>Static child</p></div>
  <div class="inline-block ..."><p>Static sibling</p></div>
  <!-- Static parent -->
  <div class="absolute ..."><p>Absolute child</p></div>
  <div class="inline-block ..."><p>Static sibling</p></div>
</div>
```

--------------------------------

### CSS accent-color property

Source: https://tailwindcss.com/docs/accent-color

The `accent-color` CSS property is used to set the accent color. This example shows the CSS generated for `accent-gray-900`.

```css
accent-color: var(--color-gray-900); /* oklch(21% 0.034 264.665) */
```

--------------------------------

### Import Tailwind CSS in app.css

Source: https://tailwindcss.com/docs/installation/framework-guides/symfony

Add Tailwind CSS import and configure source scanning to exclude the public directory and prevent watch mode recompilation loops.

```css
@import "tailwindcss";
@source not "../../public";
```

--------------------------------

### Responsive Scroll Padding

Source: https://tailwindcss.com/docs/scroll-padding

Apply scroll padding utilities at specific breakpoints using variants like `md:`. For example, `md:scroll-p-0` applies at medium screens and up.

```html
<div class="scroll-p-8 md:scroll-p-0 ...">  <!-- ... --></div>
```

--------------------------------

### justify-center and justify-center-safe

Source: https://tailwindcss.com/docs/justify-content

Use `justify-center` or `justify-center-safe` to align items along the center of the container's main axis. `justify-center-safe` aligns to the start when there isn't enough space.

```html
<div class="flex justify-center ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

```html
<div class="flex justify-center-safe ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Basic Contrast Utilities

Source: https://tailwindcss.com/docs/filter-contrast

Apply predefined contrast levels to an element using classes like contrast-50, contrast-100, contrast-125, and contrast-200.

```html
<img class="contrast-50 ..." src="/img/mountains.jpg" /><img class="contrast-100 ..." src="/img/mountains.jpg" /><img class="contrast-125 ..." src="/img/mountains.jpg" /><img class="contrast-200 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Basic Dark Mode Styling with Tailwind

Source: https://tailwindcss.com/docs/dark-mode

Use the `dark:` variant to apply styles when dark mode is enabled. This example shows a card component with different background and text colors for light and dark modes.

```html
<div class="bg-white dark:bg-gray-800 rounded-lg px-6 py-8 ring shadow-xl ring-gray-900/5">  <div>    <span class="inline-flex items-center justify-center rounded-md bg-indigo-500 p-2 shadow-lg">      <svg class="h-6 w-6 stroke-white" ...>        <!-- ... -->      </svg>    </span>  </div>  <h3 class="text-gray-900 dark:text-white mt-5 text-base font-medium tracking-tight ">Writes upside-down</h3>  <p class="text-gray-500 dark:text-gray-400 mt-2 text-sm ">    The Zero Gravity Pen can be used to write in any orientation, including upside-down. It even works in outer space.  </p></div>
```

--------------------------------

### Set Bottom Border Color with Gray Palette

Source: https://tailwindcss.com/docs/border-color

Use the `border-b-{color}-{shade}` utility to set the bottom border color. This example shows shades of gray.

```html
border-b-gray-900
border-b-gray-950
```

--------------------------------

### X-Axis Element Scaling

Source: https://tailwindcss.com/docs/scale

Utilize `scale-x-<number>` utilities to scale an element specifically on the x-axis by a percentage of its original width. Examples include `scale-x-75` and `scale-x-125`.

```html
<img class="scale-x-75 ..." src="/img/mountains.jpg" /><img class="scale-x-100 ..." src="/img/mountains.jpg" /><img class="scale-x-125 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Setting Top Border Color with Pink

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses various shades of pink.

```html
<div class="border-t-pink-50"></div>
```

```html
<div class="border-t-pink-100"></div>
```

```html
<div class="border-t-pink-200"></div>
```

```html
<div class="border-t-pink-300"></div>
```

```html
<div class="border-t-pink-400"></div>
```

```html
<div class="border-t-pink-500"></div>
```

```html
<div class="border-t-pink-600"></div>
```

```html
<div class="border-t-pink-700"></div>
```

```html
<div class="border-t-pink-800"></div>
```

```html
<div class="border-t-pink-900"></div>
```

```html
<div class="border-t-pink-950"></div>
```

--------------------------------

### Basic Mask Positioning

Source: https://tailwindcss.com/docs/mask-position

Use predefined utilities like mask-center, mask-right, and mask-left-top to control the position of an element's mask image. These examples demonstrate various positions for the mask.

```html
<div class="mask-top-left mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-top mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-top-right mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-left mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-center mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-right mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-bottom-left mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-bottom mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-bottom-right mask-[url(/img/circle.png)] mask-size-[50%] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Conditional Styling for Open/Closed State (Details Element)

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles when a `<details>` element is open. This example changes the border and background of a disclosure widget when it's expanded.

```html
<details class="border border-transparent open:border-black/10 open:bg-gray-100 ..." open>  <summary class="text-sm leading-6 font-semibold text-gray-900 select-none">Why do they call it Ovaltine?</summary>  <div class="mt-3 text-sm leading-6 text-gray-600">    <p>The mug is round. The jar is round. They should call it Roundtine.</p>  </div></details>
```

--------------------------------

### Conditional Styling with Data Attributes (Exists)

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles based on the presence of a data attribute. This example adds a border to a div only when the `data-active` attribute is present.

```html
<!-- Will apply --><div data-active class="border border-gray-300 data-active:border-purple-500">  <!-- ... --></div><!-- Will not apply --><div class="border border-gray-300 data-active:border-purple-500">  <!-- ... --></div>
```

--------------------------------

### Set Border Inline Color - Indigo

Source: https://tailwindcss.com/docs/border-color

Use `border-x-{color}-{shade}` utilities to set the inline border color. This example shows various shades of indigo.

```html
<div class="border-x-indigo-50"></div>
<div class="border-x-indigo-100"></div>
<div class="border-x-indigo-200"></div>
<div class="border-x-indigo-300"></div>
<div class="border-x-indigo-400"></div>
<div class="border-x-indigo-500"></div>
<div class="border-x-indigo-600"></div>
<div class="border-x-indigo-700"></div>
<div class="border-x-indigo-800"></div>
<div class="border-x-indigo-900"></div>
<div class="border-x-indigo-950"></div>
```

--------------------------------

### Set Opacity with Tailwind CSS

Source: https://tailwindcss.com/docs/opacity

Demonstrates how to use `opacity-<number>` utilities to control the opacity of an element.

```html
<button class="bg-indigo-500 opacity-100 ..."></button><button class="bg-indigo-500 opacity-75 ..."></button><button class="bg-indigo-500 opacity-50 ..."></button><button class="bg-indigo-500 opacity-25 ..."></button>
```

--------------------------------

### Basic flex item growth and shrink

Source: https://tailwindcss.com/docs/flex

Use `flex-<number>` utilities like `flex-1` to allow a flex item to grow and shrink as needed, ignoring its initial size.

```html
<div class="flex">  <div class="w-14 flex-none ...">01</div>  <div class="w-64 flex-1 ...">02</div>  <div class="w-32 flex-1 ...">03</div></div>
```

--------------------------------

### Tailwind CSS Backdrop Saturate Utilities

Source: https://tailwindcss.com/docs/backdrop-filter-saturate

Use utilities like `backdrop-saturate-50` and `backdrop-saturate-100` utilities to control the saturation of an element's backdrop. Examples include 50%, 125%, and 200% saturation.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-saturate-50 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-saturate-125 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-saturate-200 ..."></div></div>
```

--------------------------------

### Setting Top Border Color with Gray

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses various shades of gray.

```html
<div class="border-t-gray-50"></div>
```

```html
<div class="border-t-gray-100"></div>
```

```html
<div class="border-t-gray-200"></div>
```

```html
<div class="border-t-gray-300"></div>
```

--------------------------------

### Radial Gradient Mask From Number

Source: https://tailwindcss.com/docs/mask-image

Applies a radial gradient mask, defining the 'from' position with a number. Useful for precise starting points of the gradient.

```css
mask-image: radial-gradient(var(--tw-mask-radial-shape) var(--tw-mask-radial-size) at var(--tw-mask-radial-position), black calc(var(--spacing) * <number>), transparent var(--tw-mask-radial-to));
```

--------------------------------

### Applying box-decoration-slice and box-decoration-clone in HTML

Source: https://tailwindcss.com/docs/box-decoration-break

Use these utilities to control how properties like background, border, and padding are rendered across line breaks, either as a continuous fragment (`slice`) or distinct blocks (`clone`).

```html
<span class="box-decoration-slice bg-linear-to-r from-indigo-600 to-pink-500 px-2 text-white ...">  Hello<br />World</span><span class="box-decoration-clone bg-linear-to-r from-indigo-600 to-pink-500 px-2 text-white ...">  Hello<br />World</span>
```

--------------------------------

### Applying accent-color utility

Source: https://tailwindcss.com/docs/accent-color

Use the `accent-{color}-{shade}` utility to set the accent color for form elements. For example, `accent-lime-500` applies a specific shade of lime.

```html
<input class="accent-lime-500" type="checkbox" checked>
```

```html
<input class="accent-green-500" type="radio" checked>
```

```html
<input class="accent-emerald-500" type="range">
```

--------------------------------

### Responsive Transition Easing

Source: https://tailwindcss.com/docs/transition-timing-function

Apply transition timing function utilities conditionally based on screen size using responsive variants like `md:`. This allows for different easing behaviors at different breakpoints.

```html
<button class="ease-out md:ease-in ...">  <!-- ... --></button>
```

--------------------------------

### Sort Tailwind Classes with Prettier

Source: https://tailwindcss.com/docs/editor-setup

This example shows how Tailwind CSS classes are sorted by the official Prettier plugin. Use this to maintain a consistent class order in your HTML.

```html
<!-- Before --><button class="text-white px-4 sm:px-8 py-2 sm:py-3 bg-sky-700 hover:bg-sky-800">Submit</button><!-- After --><button class="bg-sky-700 px-4 py-2 text-white hover:bg-sky-800 sm:px-8 sm:py-3">Submit</button>
```

--------------------------------

### Setting Top Border Color with Rose

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses various shades of rose.

```html
<div class="border-t-rose-50"></div>
```

```html
<div class="border-t-rose-100"></div>
```

```html
<div class="border-t-rose-200"></div>
```

```html
<div class="border-t-rose-300"></div>
```

```html
<div class="border-t-rose-400"></div>
```

```html
<div class="border-t-rose-500"></div>
```

```html
<div class="border-t-rose-600"></div>
```

```html
<div class="border-t-rose-700"></div>
```

```html
<div class="border-t-rose-800"></div>
```

```html
<div class="border-t-rose-900"></div>
```

```html
<div class="border-t-rose-950"></div>
```

--------------------------------

### Percentage-based min-height utilities

Source: https://tailwindcss.com/docs/min-height

Apply percentage-based minimum heights using utilities like min-h-full, min-h-9/10, min-h-3/4, min-h-1/2, and min-h-1/3.

```html
<div class="min-h-full ...">min-h-full</div><div class="min-h-9/10 ...">min-h-9/10</div><div class="min-h-3/4 ...">min-h-3/4</div><div class="min-h-1/2 ...">min-h-1/2</div><div class="min-h-1/3 ...">min-h-1/3</div>
```

--------------------------------

### Controlling Horizontal Space with `-space-x-px`

Source: https://tailwindcss.com/docs/margin

Use this utility to apply a negative margin of 1px to the start of elements, creating horizontal spacing. It's useful for fine-tuning gaps between adjacent elements.

```css
& > :not(:last-child) {   --tw-space-x-reverse: 0;   margin-inline-start: calc(-1px * var(--tw-space-x-reverse));   margin-inline-end: calc(-1px * calc(1 - var(--tw-space-x-reverse))); };
```

--------------------------------

### Setting Top Border Color with Slate

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses various shades of slate.

```html
<div class="border-t-slate-50"></div>
```

```html
<div class="border-t-slate-100"></div>
```

```html
<div class="border-t-slate-200"></div>
```

```html
<div class="border-t-slate-300"></div>
```

```html
<div class="border-t-slate-400"></div>
```

```html
<div class="border-t-slate-500"></div>
```

```html
<div class="border-t-slate-600"></div>
```

```html
<div class="border-t-slate-700"></div>
```

```html
<div class="border-t-slate-800"></div>
```

```html
<div class="border-t-slate-900"></div>
```

```html
<div class="border-t-slate-950"></div>
```

--------------------------------

### Responsive Flex Basis Utilities

Source: https://tailwindcss.com/docs/flex-basis

Applies flex-basis utilities at specific breakpoints. Prefix utilities with a breakpoint variant like `md:` for medium screen sizes and above.

```html
<div class="flex flex-row">  <div class="basis-1/4 md:basis-1/3">01</div>  <div class="basis-1/4 md:basis-1/3">02</div>  <div class="basis-1/2 md:basis-1/3">03</div></div>
```

--------------------------------

### Set Border Inline Color - Violet

Source: https://tailwindcss.com/docs/border-color

Use `border-x-{color}-{shade}` utilities to set the inline border color. This example shows various shades of violet.

```html
<div class="border-x-violet-50"></div>
<div class="border-x-violet-100"></div>
<div class="border-x-violet-200"></div>
<div class="border-x-violet-300"></div>
<div class="border-x-violet-400"></div>
<div class="border-x-violet-500"></div>
<div class="border-x-violet-600"></div>
<div class="border-x-violet-700"></div>
<div class="border-x-violet-800"></div>
<div class="border-x-violet-900"></div>
<div class="border-x-violet-950"></div>
```

--------------------------------

### Set Border-block Color with Mist Palette

Source: https://tailwindcss.com/docs/border-color

The `border-block-color` property controls the border color for the block axis. This example uses a shade from the mist palette.

```css
border-block-color: var(--color-mist-950); /* oklch(14.8% 0.004 228.8) */
```

--------------------------------

### Override Border Styles for Specific Elements

Source: https://tailwindcss.com/docs/preflight

Custom CSS can override Preflight styles, for example, to remove border styles for Google Maps embeds.

```css
@layer base {
  .google-map * {
    border-style: none;
  }
}
```

--------------------------------

### Creating Table Layouts with Utilities

Source: https://tailwindcss.com/docs/display

Utilize `table`, `table-row`, `table-cell`, and related utilities to construct elements that mimic native HTML table behavior. This allows for table-like layouts using utility classes.

```html
<div class="table w-full ...">
  <div class="table-header-group ...">
    <div class="table-row">
      <div class="table-cell text-left ...">Song</div>
      <div class="table-cell text-left ...">Artist</div>
      <div class="table-cell text-left ...">Year</div>
    </div>
  </div>
  <div class="table-row-group">
    <div class="table-row">
      <div class="table-cell ...">The Sliding Mr. Bones (Next Stop, Pottersville)</div>
      <div class="table-cell ...">Malcolm Lockyer</div>
      <div class="table-cell ...">1961</div>
    </div>
    <div class="table-row">
      <div class="table-cell ...">Witchy Woman</div>
      <div class="table-cell ...">The Eagles</div>
      <div class="table-cell ...">1972</div>
    </div>
    <div class="table-row">
      <div class="table-cell ...">Shining Star</div>
      <div class="table-cell ...">Earth, Wind, and Fire</div>
      <div class="table-cell ...">1975</div>
    </div>
  </div>
</div>
```

--------------------------------

### Setting Top Border Color with Fuchsia

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses various shades of fuchsia.

```html
<div class="border-t-fuchsia-50"></div>
```

```html
<div class="border-t-fuchsia-100"></div>
```

```html
<div class="border-t-fuchsia-200"></div>
```

```html
<div class="border-t-fuchsia-300"></div>
```

```html
<div class="border-t-fuchsia-400"></div>
```

```html
<div class="border-t-fuchsia-500"></div>
```

```html
<div class="border-t-fuchsia-600"></div>
```

```html
<div class="border-t-fuchsia-700"></div>
```

```html
<div class="border-t-fuchsia-800"></div>
```

```html
<div class="border-t-fuchsia-900"></div>
```

```html
<div class="border-t-fuchsia-950"></div>
```

--------------------------------

### Remove All Transforms on an Element

Source: https://tailwindcss.com/docs/transform

The `transform-none` utility removes all applied transforms to an element. Useful for conditionally disabling transforms, for example, on different screen sizes.

```html
<div class="skew-y-3 md:transform-none">  <!-- ... --></div>
```

--------------------------------

### Using Arbitrary Container Query Sizes

Source: https://tailwindcss.com/docs/responsive-design

Utilize variants like `@min-[475px]` and `@max-[960px]` for one-off container query sizes that do not need to be added to the theme.

```html
<div class="@container">  <div class="flex flex-col @min-[475px]:flex-row">    <!-- ... -->  </div></div>
```

--------------------------------

### Set Border Inline Color - Fuchsia

Source: https://tailwindcss.com/docs/border-color

Use `border-x-{color}-{shade}` utilities to set the inline border color. This example shows various shades of fuchsia.

```html
<div class="border-x-fuchsia-50"></div>
<div class="border-x-fuchsia-100"></div>
<div class="border-x-fuchsia-200"></div>
<div class="border-x-fuchsia-300"></div>
<div class="border-x-fuchsia-400"></div>
<div class="border-x-fuchsia-500"></div>
<div class="border-x-fuchsia-600"></div>
<div class="border-x-fuchsia-700"></div>
<div class="border-x-fuchsia-800"></div>
<div class="border-x-fuchsia-900"></div>
```

--------------------------------

### Include compiled CSS and use Tailwind classes

Source: https://tailwindcss.com/docs/installation/framework-guides/laravel/vite

Ensure your compiled CSS is included in the `<head>` section of your Blade template and start using Tailwind's utility classes.

```html
<!doctype html><html>  <head>    <meta charset="utf-8" />    <meta name="viewport" content="width=device-width, initial-scale=1.0" />    @vite('resources/css/app.css')  </head>  <body>    <h1 class="text-3xl font-bold underline">      Hello world!    </h1>  </body></html>
```

--------------------------------

### Responsive Font Style

Source: https://tailwindcss.com/docs/font-style

Apply font style utilities at specific breakpoints using responsive prefixes. This example makes text italic by default and normal from the medium breakpoint upwards.

```html
<p class="italic md:not-italic ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Responsive Text Color

Source: https://tailwindcss.com/docs/text-color

Apply text color utilities at specific breakpoints using prefix variants like `md:`. This example changes color from blue to green at the medium breakpoint.

```html
<p class="text-blue-600 md:text-green-600 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Set Border Inline Color - Blue

Source: https://tailwindcss.com/docs/border-color

Use `border-x-{color}-{shade}` utilities to set the inline border color. This example shows various shades of blue.

```html
<div class="border-x-blue-700"></div>
<div class="border-x-blue-800"></div>
<div class="border-x-blue-900"></div>
<div class="border-x-blue-950"></div>
```

--------------------------------

### Conditional Styling for Open/Closed State (Popover Element)

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles when a popover element is open using the `:popover-open` pseudo-class. This example transitions the opacity of a popover when it's opened.

```html
<div>  <button popovertarget="my-popover">Open Popover</button>  <div popover id="my-popover" class="opacity-0 open:opacity-100 ...">    <!-- ... -->  </div></div>
```

--------------------------------

### Using Custom Data Attribute Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles using a custom data attribute variant. This example uses the `data-checked` variant to apply an underline style.

```html
<div data-ui="checked active" class="data-checked:underline">  <!-- ... --></div>
```

--------------------------------

### Using Custom Utility with Variants

Source: https://tailwindcss.com/docs/adding-custom-styles

Shows that custom utilities, like 'content-auto', work seamlessly with standard Tailwind variants such as 'hover', 'focus', and responsive prefixes.

```html
<div class="hover:content-auto">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive background-attachment with breakpoint variants

Source: https://tailwindcss.com/docs/background-attachment

Apply background-attachment utilities conditionally at different screen sizes using breakpoint prefixes like md:.

```html
<div class="bg-local md:bg-fixed ...">  <!-- ... --></div>
```

--------------------------------

### Y-Axis Element Scaling

Source: https://tailwindcss.com/docs/scale

Employ `scale-y-<number>` utilities to scale an element on the y-axis by a percentage of its original height. Examples include `scale-y-75` and `scale-y-125`.

```html
<img class="scale-y-75 ..." src="/img/mountains.jpg" /><img class="scale-y-100 ..." src="/img/mountains.jpg" /><img class="scale-y-125 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Block-Start Border Color Utilities

Source: https://tailwindcss.com/docs/border-color

Apply these utilities to control the block-start border color. Includes options for inherit, currentColor, transparent, and predefined color scales.

```html
border-bs-inherit
border-bs-current
border-bs-transparent
border-bs-black
border-bs-white
```

```html
border-bs-red-50
border-bs-red-100
border-bs-red-200
border-bs-red-300
border-bs-red-400
border-bs-red-500
border-bs-red-600
border-bs-red-700
border-bs-red-800
border-bs-red-900
border-bs-red-950
```

```html
border-bs-orange-50
border-bs-orange-100
border-bs-orange-200
border-bs-orange-300
border-bs-orange-400
border-bs-orange-500
border-bs-orange-600
border-bs-orange-700
border-bs-orange-800
border-bs-orange-900
border-bs-orange-950
```

```html
border-bs-amber-50
border-bs-amber-100
border-bs-amber-200
border-bs-amber-300
border-bs-amber-400
border-bs-amber-500
border-bs-amber-600
border-bs-amber-700
border-bs-amber-800
border-bs-amber-900
border-bs-amber-950
```

```html
border-bs-yellow-50
border-bs-yellow-100
border-bs-yellow-200
```

--------------------------------

### Using Native CSS Variables

Source: https://tailwindcss.com/docs/compatibility

Modern browsers support native CSS variables, which Tailwind CSS leverages internally. This example demonstrates their use for styling elements.

```css
.typography {  font-size: var(--text-base);  color: var(--color-gray-700);}
```

--------------------------------

### Percentage-based max-inline-size Utilities

Source: https://tailwindcss.com/docs/max-inline-size

Utilize `max-inline-full` or `max-inline-<fraction>` utilities for percentage-based maximum inline sizes. These are useful for creating responsive layouts where elements should occupy a specific portion of their container.

```html
<div class="inline-full max-inline-9/10 ...">max-inline-9/10</div><div class="inline-full max-inline-3/4 ...">max-inline-3/4</div><div class="inline-full max-inline-1/2 ...">max-inline-1/2</div><div class="inline-full max-inline-1/3 ...">max-inline-1/3</div>
```

--------------------------------

### Spanning columns with grid-column utilities

Source: https://tailwindcss.com/docs/grid-column

Use `col-span-<number>` utilities to make an element span a specified number of columns. For example, `col-span-2` or `col-span-4`.

```html
<div class="grid grid-cols-3 gap-4">
  <div class="...">01</div>
  <div class="...">02</div>
  <div class="...">03</div>
  <div class="col-span-2 ...">04</div>
  <div class="...">05</div>
  <div class="...">06</div>
  <div class="col-span-2 ...">07</div>
</div>
```

--------------------------------

### Customize Perspective Theme

Source: https://tailwindcss.com/docs/perspective

Customize the default perspective values by defining `--perspective-*` variables within the `@theme` directive. This example adds a `perspective-remote` utility with a value of `1800px`.

```css
@theme {  --perspective-remote: 1800px; }
```

--------------------------------

### Conditional Styling with aria-checked Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles when an ARIA attribute is true. This example shows how to change the background color based on the `aria-checked` state.

```html
<div aria-checked="true" class="bg-gray-600 aria-checked:bg-sky-700">  <!-- ... --></div>
```

--------------------------------

### Applying Responsive Background Position Utilities in HTML

Source: https://tailwindcss.com/docs/background-position

Prefix a background-position utility with a breakpoint variant like md: to apply the utility only at medium screen sizes and above.

```html
<div class="bg-center md:bg-top ...">  <!-- ... --></div>
```

--------------------------------

### Translating Bare Integers to Percentages

Source: https://tailwindcss.com/docs/adding-custom-styles

Demonstrates translating a bare integer value to a percentage using separate `--value()` declarations.

```css
@utility opacity-* {
  opacity: --value([percentage]);
  opacity: calc(--value(integer) * 1%);
  opacity: --value(--opacity-*);
}
```

--------------------------------

### Radial Gradient Mask From Percentage

Source: https://tailwindcss.com/docs/mask-image

Applies a radial gradient mask, defining the 'from' position with a percentage. Offers fine-grained control over the gradient's start.

```css
mask-image: radial-gradient(var(--tw-mask-radial-shape) var(--tw-mask-radial-size) at var(--tw-mask-radial-position), black <percentage>, transparent var(--tw-mask-radial-to));
```

--------------------------------

### Define a Custom Variant with Media Query and Hover State

Source: https://tailwindcss.com/docs/adding-custom-styles

Create a custom variant that applies styles conditionally, such as when a hover interaction is supported. This example uses `@media (any-hover: hover)` and `:hover`.

```css
@custom-variant any-hover {
  @media (any-hover: hover) {
    &:hover {
      @slot;
    }
  }
}
```

--------------------------------

### Apply Left Mask Gradient From Color

Source: https://tailwindcss.com/docs/mask-image

Use `mask-l-from-<color>` to apply a left mask gradient. The gradient starts from the specified color and transitions to transparent.

```css
mask-image: linear-gradient(to left, <color> var(--tw-mask-left-from), transparent var(--tw-mask-left-to));
```

--------------------------------

### Apply responsive overflow-wrap utilities

Source: https://tailwindcss.com/docs/overflow-wrap

Prefix `overflow-wrap` utilities with breakpoint variants like `md:` to apply them conditionally based on screen size.

```html
<p class="wrap-normal md:wrap-break-word ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Apply Bottom Mask Gradient from Color

Source: https://tailwindcss.com/docs/mask-image

Use `mask-b-from-<color>` to apply a bottom mask gradient. The gradient starts from the specified color and transitions to transparent.

```css
mask-image: linear-gradient(to bottom, <color> var(--tw-mask-bottom-from), transparent var(--tw-mask-bottom-to));
```

--------------------------------

### Responsive grid column design with breakpoints

Source: https://tailwindcss.com/docs/grid-column

Apply `grid-column`, `grid-column-start`, and `grid-column-end` utilities responsively by prefixing them with a breakpoint variant, such as `md:`.

```html
<div class="col-span-2 md:col-span-6 ...">
  <!-- ... -->
</div>
```

--------------------------------

### justify-end and justify-end-safe

Source: https://tailwindcss.com/docs/justify-content

Use `justify-end` or `justify-end-safe` to justify items against the end of the container's main axis. `justify-end-safe` aligns to the start when there isn't enough space.

```html
<div class="flex justify-end ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>03</div></div>
```

```html
<div class="flex justify-end-safe ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>03</div></div>
```

--------------------------------

### Apply custom padding with arbitrary values

Source: https://tailwindcss.com/docs/padding

Use utilities like `p-[<value>]` to set the padding based on a completely custom value.

```html
<div class="p-[5px] ...">  <!-- ... --></div>
```

--------------------------------

### Conditional Styling with Data Attributes (Specific Value)

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles based on a specific value of a data attribute. This example applies padding only when the `data-size` attribute is set to 'large'.

```html
<!-- Will apply --><div data-size="large" class="data-[size=large]:p-8">  <!-- ... --></div><!-- Will not apply --><div data-size="medium" class="data-[size=large]:p-8">  <!-- ... --></div>
```

--------------------------------

### Setting Block Border Color (Indigo)

Source: https://tailwindcss.com/docs/border-color

Use the `border-block-indigo-{shade}` utility to set the color of the block borders. For example, `border-block-indigo-900` sets the block borders to a dark indigo.

```html
<div class="border-block-indigo-900">...</div>
```

--------------------------------

### Configure PostCSS with Tailwind plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/angular

Create a .postcssrc.json configuration file in the project root to register the @tailwindcss/postcss plugin.

```json
{  "plugins": {    "@tailwindcss/postcss": {}  }}
```

--------------------------------

### Define Grid Row Start and End Lines

Source: https://tailwindcss.com/docs/grid-row

Apply `row-start-<number>` or `row-end-<number>` utilities to position an element at a specific grid line. These can be combined with `row-span-<number>` for more complex layouts.

```html
<div class="grid grid-flow-col grid-rows-3 gap-4">  <div class="row-span-2 row-start-2 ...">01</div>  <div class="row-span-2 row-end-3 ...">02</div>  <div class="row-start-1 row-end-4 ...">03</div></div>
```

--------------------------------

### Setting Block Border Color (Violet)

Source: https://tailwindcss.com/docs/border-color

Use the `border-block-violet-{shade}` utility to set the color of the block borders. For example, `border-block-violet-900` sets the block borders to a dark violet.

```html
<div class="border-block-violet-900">...</div>
```

--------------------------------

### Conditional Hover Variant in v4

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, the `hover` variant only applies if the primary input device supports hover. This CSS example shows the default behavior.

```css
@media (hover: hover) {
  .hover\:underline:hover {
    text-decoration: underline;
  }
}
```

--------------------------------

### Apply Custom Shadow Color in Markup

Source: https://tailwindcss.com/docs/text-shadow

Use the `text-shadow-regal-blue` class in your HTML to apply the text shadow with the custom color. This example applies the shadow to a paragraph.

```html
<p class="text-shadow-regal-blue">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Set Columns by Width

Source: https://tailwindcss.com/docs/columns

Use utilities like `columns-xs` to set an ideal column width. The number of columns adjusts to prevent them from becoming too narrow.

```html
<div class="columns-3xs ...">
  <img class="aspect-3/2 ..." src="/img/mountains-1.jpg" />
  <img class="aspect-square ..." src="/img/mountains-2.jpg" />
  <img class="aspect-square ..." src="/img/mountains-3.jpg" />
  <!-- ... -->
</div>
```

--------------------------------

### Reset and Define All Breakpoints

Source: https://tailwindcss.com/docs/responsive-design

Completely replace all default breakpoints by resetting them to `initial` and then defining your custom set from scratch.

```css
@import "tailwindcss";@theme {  --breakpoint-*: initial;  --breakpoint-tablet: 40rem;  --breakpoint-laptop: 64rem;  --breakpoint-desktop: 80rem;}
```

--------------------------------

### Apply Custom Text Shadow in Markup

Source: https://tailwindcss.com/docs/text-shadow

Use the `text-shadow-xl` class in your HTML to apply the customized text shadow. This example shows how to apply it to a paragraph element.

```html
<p class="text-shadow-xl">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Container scale min-width

Source: https://tailwindcss.com/docs/min-width

Apply fixed minimum widths using container scale utilities like min-w-sm, min-w-xl, etc., for responsive container-based sizing.

```html
<div class="w-40 ...">
  <div class="min-w-lg ...">min-w-lg</div>
  <div class="min-w-md ...">min-w-md</div>
  <div class="min-w-sm ...">min-w-sm</div>
  <div class="min-w-xs ...">min-w-xs</div>
  <div class="min-w-2xs ...">min-w-2xs</div>
  <div class="min-w-3xs ...">min-w-3xs</div>
</div>
```

--------------------------------

### Negative Rotation with Tailwind CSS

Source: https://tailwindcss.com/docs/rotate

Use `-rotate-<number>` utilities like `-rotate-45` and `-rotate-90` to rotate an element counterclockwise by degrees. Examples include -45, -90, and -210 degrees.

```html
<img class="-rotate-45 ..." src="/img/mountains.jpg" /><img class="-rotate-90 ..." src="/img/mountains.jpg" /><img class="-rotate-210 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Create a new Ember.js project

Source: https://tailwindcss.com/docs/installation/framework-guides/emberjs

Use Ember CLI to generate a new Ember.js application with Embroider enabled and without the welcome page.

```bash
npx ember-cli new my-project --embroider --no-welcomecd my-project
```

--------------------------------

### Responsive min-width with breakpoint variants

Source: https://tailwindcss.com/docs/min-width

Prefix min-width utilities with breakpoint variants like md: to apply the utility only at specific screen sizes.

```html
<div class="w-24 min-w-full md:min-w-0 ...">
  <!-- ... -->
</div>
```

--------------------------------

### CSS variable resolution fallback example

Source: https://tailwindcss.com/docs/theme

Demonstrates how CSS variables resolve at their definition point, causing fallback to sans-serif instead of Inter when --font-inter is defined deeper in the DOM tree.

```HTML
<div id="parent" style="--font-sans: var(--font-inter, sans-serif);">
  <div id="child" style="--font-inter: Inter; font-family: var(--font-sans);">
    This text will use the sans-serif font, not Inter.
  </div>
</div>
```

--------------------------------

### JavaScript for System Theme and Manual Toggling

Source: https://tailwindcss.com/docs/dark-mode

Use JavaScript to manage dark mode preferences, respecting the system's `prefers-color-scheme` and allowing manual toggling via `localStorage`.

```javascript
// On page load or when changing themes, best to add inline in `head` to avoid FOUCdocument.documentElement.classList.toggle(  "dark",  localStorage.theme === "dark" ||    (!("theme" in localStorage) && window.matchMedia("(prefers-color-scheme: dark)").matches),);// Whenever the user explicitly chooses light modelocalStorage.theme = "light";
// Whenever the user explicitly chooses dark modelocalStorage.theme = "dark";
// Whenever the user explicitly chooses to respect the OS preferencelocalStorage.removeItem("theme");
```

--------------------------------

### Responsive mask size

Source: https://tailwindcss.com/docs/mask-size

Prefix a `mask-size` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example applies `mask-auto` by default and `mask-contain` from the medium breakpoint onwards.

```html
<div class="mask-auto md:mask-contain ...">  <!-- ... --></div>
```

--------------------------------

### Setting Top Block Border Color (Violet)

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-violet-{shade}` utility to set the color of the top border. For example, `border-t-violet-950` sets the top border to a very dark violet.

```html
<div class="border-t-violet-950">...</div>
```

--------------------------------

### Responsive align-items

Source: https://tailwindcss.com/docs/align-items

Apply `align-items` utilities at different screen sizes using responsive prefixes like `md:`.

```html
<div class="flex items-stretch md:items-center ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive background-clip utilities

Source: https://tailwindcss.com/docs/background-clip

Apply `background-clip` utilities at specific breakpoints using responsive variants like `md:`.

```html
<div class="bg-clip-border md:bg-clip-padding ...">  <!-- ... --></div>
```

--------------------------------

### mask-contain utility

Source: https://tailwindcss.com/docs/mask-size

Use the `mask-contain` utility to scale the mask image to the outer edges without cropping or stretching. This example also sets a background image.

```html
<div class="mask-contain mask-[url(/img/scribble.png)] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Setting Top Block Border Color (Purple)

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-purple-{shade}` utility to set the color of the top border. For example, `border-t-purple-950` sets the top border to a very dark purple.

```html
<div class="border-t-purple-950">...</div>
```

--------------------------------

### Setting Top Block Border Color (Fuchsia)

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-fuchsia-{shade}` utility to set the color of the top border. For example, `border-t-fuchsia-950` sets the top border to a very dark fuchsia.

```html
<div class="border-t-fuchsia-950">...</div>
```

--------------------------------

### Setting Top Block Border Color (Indigo)

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-indigo-{shade}` utility to set the color of the top border. For example, `border-t-indigo-950` sets the top border to a very dark indigo.

```html
<div class="border-t-indigo-950">...</div>
```

--------------------------------

### Relative Positioning Example

Source: https://tailwindcss.com/docs/position

Use the `relative` utility to position an element according to the normal flow. Offsets are relative to the element's normal position, and it acts as a position reference for absolutely positioned children.

```html
<div class="relative ...">
  <p>Relative parent</p>
  <div class="absolute bottom-0 left-0 ...">
    <p>Absolute child</p>
  </div>
</div>
```

--------------------------------

### Basic Mask Mode Application with mask-alpha and mask-luminance (HTML)

Source: https://tailwindcss.com/docs/mask-mode

This snippet demonstrates applying `mask-alpha` and `mask-luminance` utilities to control an element's mask mode. `mask-alpha` uses opacity, while `mask-luminance` uses grayscale luminance for visibility.

```html
<div class="mask-alpha mask-r-from-black mask-r-from-50% mask-r-to-transparent bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-luminance mask-r-from-white mask-r-from-50% mask-r-to-black bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Responsive max-inline-size Utilities

Source: https://tailwindcss.com/docs/max-inline-size

Control maximum inline size responsively by prefixing utilities with breakpoint variants like `md:`. This ensures the utility is applied only at the specified screen size and above.

```html
<div class="max-inline-sm md:max-inline-lg ...">  <!-- ... --></div>
```

--------------------------------

### Configure @tailwindcss/postcss in postcss.config.js

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Create a `postcss.config.js` file in the root of your project and add `@tailwindcss/postcss` to the plugins object.

```JavaScript
module.exports = {  plugins: {    "@tailwindcss/postcss": {},  },};
```

--------------------------------

### Responsive Grayscale Design

Source: https://tailwindcss.com/docs/filter-grayscale

Prefix a `filter: grayscale()` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example shows applying grayscale on larger screens while it's off on smaller ones.

```html
<img class="grayscale md:grayscale-0 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Alternative to pseudo-elements using a span

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

For many cases, using a regular HTML element like `<span>` is simpler and results in less code than using `::before` or `::after` pseudo-elements. This example shows the same design using a span.

```html
<blockquote class="text-center text-2xl font-semibold text-gray-900 italic">
  When you look  <span class="relative">
    <span class="absolute -inset-1 block -skew-y-3 bg-pink-500" aria-hidden="true"></span>
    <span class="relative text-white">annoyed</span>
  </span>  all the time, people think that you're busy.
</blockquote>
```

--------------------------------

### Isolating Blending with `isolate`

Source: https://tailwindcss.com/docs/mix-blend-mode

Use the `isolate` utility on a parent element to create a new stacking context, preventing blending with content behind it. This example demonstrates blending with and without isolation.

```html
<div class="isolate flex justify-center -space-x-14">  <div class="bg-yellow-500 mix-blend-multiply ..."></div>  <div class="bg-green-500 mix-blend-multiply ..."></div></div><div class="flex justify-center -space-x-14">  <div class="bg-yellow-500 mix-blend-multiply ..."></div>  <div class="bg-green-500 mix-blend-multiply ..."></div></div>
```

--------------------------------

### Responsive transition duration with breakpoint variant

Source: https://tailwindcss.com/docs/transition-duration

Prefix duration utilities with breakpoint variants like md: to apply different transition durations at specific screen sizes.

```html
<button class="duration-0 md:duration-150 ...">  <!-- ... --></button>
```

--------------------------------

### Apply Left Mask Gradient From Value

Source: https://tailwindcss.com/docs/mask-image

Use `mask-l-from-[<value>]` to apply a left mask gradient. The gradient starts from black at the specified value and transitions to transparent.

```css
mask-image: linear-gradient(to left, black <value>, transparent var(--tw-mask-left-to));
```

--------------------------------

### Apply Left Mask Gradient From Percentage

Source: https://tailwindcss.com/docs/mask-image

Use `mask-l-from-<percentage>` to apply a left mask gradient. The gradient starts from black at the specified percentage and transitions to transparent.

```css
mask-image: linear-gradient(to left, black <percentage>, transparent var(--tw-mask-left-to));
```

--------------------------------

### Responsive flex-shrink utility

Source: https://tailwindcss.com/docs/flex-shrink

Apply `flex-shrink` utilities at specific breakpoints using responsive variants like `md:`. This allows for adaptive layouts based on screen size.

```html
<div class="shrink md:shrink-0 ...">  <!-- ... --></div>
```

--------------------------------

### Apply Bottom Mask Gradient from Value

Source: https://tailwindcss.com/docs/mask-image

Use `mask-b-from-[<value>]` to apply a bottom mask gradient. The gradient starts from black at a specified value and transitions to transparent.

```css
mask-image: linear-gradient(to bottom, black <value>, transparent var(--tw-mask-bottom-to));
```

--------------------------------

### Basic touch-action utilities in HTML

Source: https://tailwindcss.com/docs/touch-action

Apply `touch-action` utilities to control panning and zooming behavior on elements within an `overflow-auto` container.

```html
<div class="h-48 w-full touch-auto overflow-auto ...">  <img class="h-auto w-[150%] max-w-none" src="..." /></div><div class="h-48 w-full touch-none overflow-auto ...">  <img class="h-auto w-[150%] max-w-none" src="..." /></div><div class="h-48 w-full touch-pan-x overflow-auto ...">  <img class="h-auto w-[150%] max-w-none" src="..." /></div><div class="h-48 w-full touch-pan-y overflow-auto ...">  <img class="h-auto w-[150%] max-w-none" src="..." /></div>
```

--------------------------------

### Responsive Positioning Utilities

Source: https://tailwindcss.com/docs/top-right-bottom-left

Prefix positioning utilities with responsive variants like `md:` to apply styles only at specific breakpoints and above.

```html
<div class="top-4 md:top-6 ...">  <!-- ... --></div>
```

--------------------------------

### Apply Bottom Mask Gradient from Percentage

Source: https://tailwindcss.com/docs/mask-image

Use `mask-b-from-<percentage>` to apply a bottom mask gradient. The gradient starts from black at the specified percentage and transitions to transparent.

```css
mask-image: linear-gradient(to bottom, black <percentage>, transparent var(--tw-mask-bottom-to));
```

--------------------------------

### Set border-y color to amber shades

Source: https://tailwindcss.com/docs/border-color

Apply different shades of amber to the top and bottom borders using the `border-y-{color}-{shade}` utility. Examples range from light to dark amber.

```html
<div class="border-y-amber-50">...</div>
```

```html
<div class="border-y-amber-100">...</div>
```

```html
<div class="border-y-amber-200">...</div>
```

```html
<div class="border-y-amber-300">...</div>
```

```html
<div class="border-y-amber-400">...</div>
```

```html
<div class="border-y-amber-500">...</div>
```

```html
<div class="border-y-amber-600">...</div>
```

```html
<div class="border-y-amber-700">...</div>
```

```html
<div class="border-y-amber-800">...</div>
```

```html
<div class="border-y-amber-900">...</div>
```

```html
<div class="border-y-amber-950">...</div>
```

--------------------------------

### Responsive Background Size

Source: https://tailwindcss.com/docs/background-size

Apply background size utilities responsively by prefixing them with a breakpoint variant, such as `md:`. This allows the background size to change at different screen sizes.

```html
<div class="bg-auto md:bg-contain ...">  <!-- ... --></div>
```

--------------------------------

### Responsive design for numeric font variants

Source: https://tailwindcss.com/docs/font-variant-numeric

Prefix a `font-variant-numeric` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This example switches from proportional to tabular figures at the `md` breakpoint.

```html
<p class="proportional-nums md:tabular-nums ...">
  Lorem ipsum dolor sit amet...
</p>
```

--------------------------------

### Configure Prefix for Utilities and Variables

Source: https://tailwindcss.com/docs/preflight

Apply a prefix like `tw` to utilities and theme variables by adding `prefix(tw)` to both the `theme.css` and `utilities.css` imports when importing Tailwind CSS files individually.

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/utilities.css" layer(utilities);
@import "tailwindcss/theme.css" layer(theme) prefix(tw);
@import "tailwindcss/utilities.css" layer(utilities) prefix(tw);
```

--------------------------------

### Apply responsive place-content utilities in Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Prefix `place-content` utilities with a breakpoint variant like `md:` to apply them conditionally based on screen size.

```html
<div class="grid place-content-start md:place-content-center ...">  <!-- ... --></div>
```

--------------------------------

### Container scale widths

Source: https://tailwindcss.com/docs/width

Use utilities like w-sm and w-xl to set fixed widths based on the container scale.

```html
<div class="w-xl ...">w-xl</div><div class="w-lg ...">w-lg</div><div class="w-md ...">w-md</div><div class="w-sm ...">w-sm</div><div class="w-xs ...">w-xs</div><div class="w-2xs ...">w-2xs</div><div class="w-3xs ...">w-3xs</div>
```

--------------------------------

### Background Size: Auto

Source: https://tailwindcss.com/docs/background-size

Use `bg-auto` to display the background image at its intrinsic size. Combine with `bg-no-repeat` to prevent tiling.

```html
<div class="bg-[url(/img/mountains.jpg)] bg-auto bg-center bg-no-repeat"></div>
```

--------------------------------

### Setting Top Border Color with Purple

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the top border color. This example uses a dark purple shade.

```html
<div class="border-t-purple-950"></div>
```

--------------------------------

### Set Percentage-Based Block Size

Source: https://tailwindcss.com/docs/block-size

Use `block-full` or `block-<fraction>` utilities like `block-1/2` and `block-2/5` to give an element a percentage-based block size.

```html
<div class="block-full ...">block-full</div><div class="block-9/10 ...">block-9/10</div><div class="block-3/4 ...">block-3/4</div><div class="block-1/2 ...">block-1/2</div><div class="block-1/3 ...">block-1/3</div>
```

--------------------------------

### Masking Horizontal Edges with Start and End

Source: https://tailwindcss.com/docs/mask-image

Apply a linear gradient mask to both the left and right edges simultaneously using `mask-x-from-<value>` and `mask-x-to-<value>`. The background image is also applied.

```html
<div class="mask-x-from-70% mask-x-to-90% bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Masking Vertical Edges with Start and End

Source: https://tailwindcss.com/docs/mask-image

Apply a linear gradient mask to both the top and bottom edges simultaneously using `mask-y-from-<value>` and `mask-y-to-<value>`. The background image is also applied.

```html
<div class="mask-y-from-70% mask-y-to-90% bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Responsive Appearance

Source: https://tailwindcss.com/docs/appearance

Prefix an `appearance` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<select class="appearance-auto md:appearance-none ...">
  <!-- ... -->
</select>
```

--------------------------------

### Set Border-y Color with Olive Palette

Source: https://tailwindcss.com/docs/border-color

Use `border-y-{color}-{shade}` utilities to set the vertical border color. This example uses a shade from the olive palette.

```html
<div class="border-y-olive-950">...</div>
```

--------------------------------

### Custom brightness value with arbitrary syntax

Source: https://tailwindcss.com/docs/filter-brightness

Use the brightness-[<value>] syntax to apply a custom brightness value not covered by standard utility classes.

```html
<img class="brightness-[1.75] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Set border-block-color to orange-950

Source: https://tailwindcss.com/docs/border-color

The `border-block-color` CSS property can be used to set the color of borders in the block direction. This example uses a CSS variable for a dark orange shade.

```css
.border-y-orange-950 { border-block-color: var(--color-orange-950); /* oklch(26.6% 0.079 36.259) */ }
```

--------------------------------

### Configure PostCSS with Tailwind plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/nextjs

Create a postcss.config.mjs file in the project root and register the @tailwindcss/postcss plugin.

```javascript
const config = {
  plugins: {
    "@tailwindcss/postcss": {},
  },
};

export default config;
```

--------------------------------

### Set border-y color to orange-950

Source: https://tailwindcss.com/docs/border-color

Use the `border-y-{color}-{shade}` utility to set the top and bottom border color. This example uses a dark orange shade.

```html
<div class="border-y-orange-950">...</div>
```

--------------------------------

### Responsive Box Shadows

Source: https://tailwindcss.com/docs/box-shadow

Apply box shadow utilities responsively by prefixing them with a breakpoint variant, like `md:shadow-lg`. This ensures the shadow is only applied at medium screen sizes and above.

```html
<div class="shadow-none md:shadow-lg ...">  <!-- ... --></div>
```

--------------------------------

### Custom line clamping with line-clamp-(<custom-property>)

Source: https://tailwindcss.com/docs/line-clamp

Use the `line-clamp-(<custom-property>)` syntax as a shorthand for `line-clamp-[var(<custom-property>)]` to automatically include the `var()` function.

```html
<p class="line-clamp-(--my-line-count) ...">
  Lorem ipsum dolor sit amet...
</p>
```

--------------------------------

### Set Border Inline Color - Purple

Source: https://tailwindcss.com/docs/border-color

Use `border-x-{color}-{shade}` utilities to set the inline border color. This example shows various shades of purple.

```html
<div class="border-x-purple-50"></div>
<div class="border-x-purple-100"></div>
<div class="border-x-purple-200"></div>
<div class="border-x-purple-300"></div>
<div class="border-x-purple-400"></div>
<div class="border-x-purple-500"></div>
<div class="border-x-purple-600"></div>
<div class="border-x-purple-700"></div>
<div class="border-x-purple-800"></div>
<div class="border-x-purple-900"></div>
<div class="border-x-purple-950"></div>
```

--------------------------------

### Conditional Styling for RTL/LTR Layouts

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles conditionally based on the text direction (RTL or LTR). This example uses `ltr:ml-3` and `rtl:mr-3` for margin adjustments in a user profile card.

```html
<div class="group flex items-center">  <img class="h-12 w-12 shrink-0 rounded-full" src="..." alt="" />  <div class="ltr:ml-3 rtl:mr-3">    <p class="text-gray-700 group-hover:text-gray-900 ...">...</p>    <p class="text-gray-500 group-hover:text-gray-700 ...">...</p>  </div></div>
```

--------------------------------

### Import app.css into SvelteKit layout

Source: https://tailwindcss.com/docs/installation/framework-guides/sveltekit

Imports the `app.css` file into the main SvelteKit layout component (`./src/routes/+layout.svelte`) to apply global styles.

```svelte
<script>
  let { children } = $props();
  import "../app.css";
</script>{@render children()}
```

--------------------------------

### Basic list-style-type utilities

Source: https://tailwindcss.com/docs/list-style-type

Apply list marker styles using list-disc, list-decimal, or list-none classes on ul or ol elements.

```html
<ul class="list-disc">
  <li>Now this is a story all about how, my life got flipped-turned upside down</li>
  <!-- ... -->
</ul>
<ol class="list-decimal">
  <li>Now this is a story all about how, my life got flipped-turned upside down</li>
  <!-- ... -->
</ol>
<ul class="list-none">
  <li>Now this is a story all about how, my life got flipped-turned upside down</li>
  <!-- ... -->
</ul>
```

--------------------------------

### Basic background-origin utilities

Source: https://tailwindcss.com/docs/background-origin

Use the `bg-origin-border`, `bg-origin-padding`, and `bg-origin-content` utilities to control where an element's background is rendered. These utilities correspond to `background-origin: border-box;`, `background-origin: padding-box;`, and `background-origin: content-box;` respectively.

```html
<div class="border-4 bg-[url(/img/mountains.jpg)] bg-origin-border p-3 ..."></div><div class="border-4 bg-[url(/img/mountains.jpg)] bg-origin-padding p-3 ..."></div><div class="border-4 bg-[url(/img/mountains.jpg)] bg-origin-content p-3 ..."></div>
```

--------------------------------

### Using Arbitrary Values for Positioning

Source: https://tailwindcss.com/docs/adding-custom-styles

Generate utility classes on the fly with arbitrary values for precise styling, such as setting a specific pixel value for 'top'.

```html
<div class="top-[117px]">  <!-- ... --></div>
```

--------------------------------

### Responsive White-space Design

Source: https://tailwindcss.com/docs/white-space

Prefix a white-space utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<p class="whitespace-pre md:whitespace-normal ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Responsive Backface Visibility in Tailwind CSS

Source: https://tailwindcss.com/docs/backface-visibility

Apply backface visibility utilities conditionally based on screen size using responsive variants. This example shows `backface-visible` by default, switching to `backface-hidden` at medium screen sizes and above.

```html
<div class="backface-visible md:backface-hidden ...">  <!-- ... --></div>
```

--------------------------------

### Basic Drop Shadows

Source: https://tailwindcss.com/docs/filter-drop-shadow

Apply predefined drop shadow sizes like sm, md, lg, xl to elements using utility classes.

```html
<svg class="drop-shadow-md ...">  <!-- ... --></svg><svg class="drop-shadow-lg ...">  <!-- ... --></svg><svg class="drop-shadow-xl ...">  <!-- ... --></svg>
```

--------------------------------

### Apply padding to individual sides with pt/pr/pb/pl-<number>

Source: https://tailwindcss.com/docs/padding

Use `pt-<number>`, `pr-<number>`, `pb-<number>`, and `pl-<number>` utilities to control padding on specific sides of an element.

```html
<div class="pt-6 ...">pt-6</div><div class="pr-4 ...">pr-4</div><div class="pb-8 ...">pb-8</div><div class="pl-2 ...">pl-2</div>
```

--------------------------------

### Arbitrary variant for custom selectors

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use arbitrary variants by wrapping a custom selector string in square brackets. This example changes the cursor when an element has the `is-dragging` class.

```html
<ul role="list">
  {#each items as item}
    <li class="[&.is-dragging]:cursor-grabbing">{item}</li>
  {/each}
</ul>
```

--------------------------------

### Tailwind CSS Text Color Utilities

Source: https://tailwindcss.com/docs/text-color

These examples demonstrate the usage of Tailwind CSS text color utility classes. Each class maps to a specific CSS variable representing a color in the Tailwind color palette.

```css
text-lime-950 { color: var(--color-lime-950); /* oklch(27.4% 0.072 132.109) */ }
```

```css
text-green-50 { color: var(--color-green-50); /* oklch(98.2% 0.018 155.826) */ }
```

```css
text-green-100 { color: var(--color-green-100); /* oklch(96.2% 0.044 156.743) */ }
```

```css
text-green-200 { color: var(--color-green-200); /* oklch(92.5% 0.084 155.995) */ }
```

```css
text-green-300 { color: var(--color-green-300); /* oklch(87.1% 0.15 154.449) */ }
```

```css
text-green-400 { color: var(--color-green-400); /* oklch(79.2% 0.209 151.711) */ }
```

```css
text-green-500 { color: var(--color-green-500); /* oklch(72.3% 0.219 149.579) */ }
```

```css
text-green-600 { color: var(--color-green-600); /* oklch(62.7% 0.194 149.214) */ }
```

```css
text-green-700 { color: var(--color-green-700); /* oklch(52.7% 0.154 150.069) */ }
```

```css
text-green-800 { color: var(--color-green-800); /* oklch(44.8% 0.119 151.328) */ }
```

```css
text-green-900 { color: var(--color-green-900); /* oklch(39.3% 0.095 152.535) */ }
```

```css
text-green-950 { color: var(--color-green-950); /* oklch(26.6% 0.065 152.934) */ }
```

```css
text-emerald-50 { color: var(--color-emerald-50); /* oklch(97.9% 0.021 166.113) */ }
```

```css
text-emerald-100 { color: var(--color-emerald-100); /* oklch(95% 0.052 163.051) */ }
```

```css
text-emerald-200 { color: var(--color-emerald-200); /* oklch(90.5% 0.093 164.15) */ }
```

```css
text-emerald-300 { color: var(--color-emerald-300); /* oklch(84.5% 0.143 164.978) */ }
```

```css
text-emerald-400 { color: var(--color-emerald-400); /* oklch(76.5% 0.177 163.223) */ }
```

```css
text-emerald-500 { color: var(--color-emerald-500); /* oklch(69.6% 0.17 162.48) */ }
```

```css
text-emerald-600 { color: var(--color-emerald-600); /* oklch(59.6% 0.145 163.225) */ }
```

```css
text-emerald-700 { color: var(--color-emerald-700); /* oklch(50.8% 0.118 165.612) */ }
```

```css
text-emerald-800 { color: var(--color-emerald-800); /* oklch(43.2% 0.095 166.913) */ }
```

```css
text-emerald-900 { color: var(--color-emerald-900); /* oklch(37.8% 0.077 168.94) */ }
```

```css
text-emerald-950 { color: var(--color-emerald-950); /* oklch(26.2% 0.051 172.552) */ }
```

```css
text-teal-50 { color: var(--color-teal-50); /* oklch(98.4% 0.014 180.72) */ }
```

```css
text-teal-100 { color: var(--color-teal-100); /* oklch(95.3% 0.051 180.801) */ }
```

```css
text-teal-200 { color: var(--color-teal-200); /* oklch(91% 0.096 180.426) */ }
```

```css
text-teal-300 { color: var(--color-teal-300); /* oklch(85.5% 0.138 181.071) */ }
```

```css
text-teal-400 { color: var(--color-teal-400); /* oklch(77.7% 0.152 181.912) */ }
```

```css
text-teal-500 { color: var(--color-teal-500); /* oklch(70.4% 0.14 182.503) */ }
```

```css
text-teal-600 { color: var(--color-teal-600); /* oklch(60% 0.118 184.704) */ }
```

```css
text-teal-700 { color: var(--color-teal-700); /* oklch(51.1% 0.096 186.391) */ }
```

```css
text-teal-800 { color: var(--color-teal-800); /* oklch(43.7% 0.078 188.216) */ }
```

```css
text-teal-900 { color: var(--color-teal-900); /* oklch(38.6% 0.063 188.416) */ }
```

```css
text-teal-950 { color: var(--color-teal-950); /* oklch(27.7% 0.046 192.524) */ }
```

```css
text-cyan-50 { color: var(--color-cyan-50); /* oklch(98.4% 0.019 200.873) */ }
```

```css
text-cyan-100 { color: var(--color-cyan-100); /* oklch(95.6% 0.045 203.388) */ }
```

```css
text-cyan-200 { color: var(--color-cyan-200); /* oklch(91.7% 0.08 205.041) */ }
```

```css
text-cyan-300 { color: var(--color-cyan-300); /* oklch(86.5% 0.127 207.078) */ }
```

```css
text-cyan-400 { color: var(--color-cyan-400); /* oklch(78.9% 0.154 211.53) */ }
```

```css
text-cyan-500 { color: var(--color-cyan-500); /* oklch(71.5% 0.143 215.221) */ }
```

```css
text-cyan-600 { color: var(--color-cyan-600); /* oklch(60.9% 0.126 221.723) */ }
```

```css
text-cyan-700 { color: var(--color-cyan-700); /* oklch(52% 0.105 223.128) */ }
```

```css
text-cyan-800 { color: var(--color-cyan-800); /* oklch(45% 0.085 224.283) */ }
```

```css
text-cyan-900 { color: var(--color-cyan-900); /* oklch(39.8% 0.07 227.392) */ }
```

```css
text-cyan-950 { color: var(--color-cyan-950); /* oklch(30.2% 0.056 229.695) */ }
```

```css
text-sky-50 { color: var(--color-sky-50); /* oklch(97.7% 0.013 236.62) */ }
```

```css
text-sky-100 { color: var(--color-sky-100); /* oklch(95.1% 0.026 236.824) */ }
```

```css
text-sky-200 { color: var(--color-sky-200); /* oklch(90.1% 0.058 230.902) */ }
```

```css
text-sky-300 { color: var(--color-sky-300); /* oklch(82.8% 0.111 230.318) */ }
```

```css
text-sky-400 { color: var(--color-sky-400); /* oklch(74.6% 0.16 232.661) */ }
```

```css
text-sky-500 { color: var(--color-sky-500); /* oklch(68.5% 0.169 237.323) */ }
```

```css
text-sky-600 { color: var(--color-sky-600); /* oklch(58.8% 0.158 241.966) */ }
```

```css
text-sky-700 { color: var(--color-sky-700); /* oklch(50% 0.134 242.749) */ }
```

```css
text-sky-800 { color: var(--color-sky-800); /* oklch(44.3% 0.11 240.79) */ }
```

```css
text-sky-900 { color: var(--color-sky-900); /* oklch(39.1% 0.09 240.876) */ }
```

```css
text-sky-950 { color: var(--color-sky-950); /* oklch(29.3% 0.066 243.157) */ }
```

```css
text-blue-50 { color: var(--color-blue-50); /* oklch(97% 0.014 254.604) */ }
```

```css
text-blue-100 { color: var(--color-blue-100); /* oklch(93.2% 0.032 255.585) */ }
```

```css
text-blue-200 { color: var(--color-blue-200); /* oklch(88.2% 0.059 254.128) */ }
```

--------------------------------

### Apply Custom Background Color Utility

Source: https://tailwindcss.com/docs/background-color

After defining a custom background color variable, you can use the corresponding utility class in your HTML markup. For example, `bg-regal-blue` will apply the color defined by `--color-regal-blue`.

```html
<div class="bg-regal-blue">  <!-- ... --></div>
```

--------------------------------

### Set Border Width Using Logical Properties (Block Start/End)

Source: https://tailwindcss.com/docs/border-width

Use `border-bs-*` and `border-be-*` for `border-block-start-width` and `border-block-end-width`. These properties adjust based on the writing mode.

```html
<div class="border-bs-4 ..."></div>
```

--------------------------------

### Compiled CSS with 'prefix' option

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Demonstrates compiled CSS when the 'prefix' option is used, showing prefixed utility classes and CSS variables.

```css
@layer theme {
  :root {
    --tw-color-red-500: oklch(0.637 0.237 25.331);
  }
}

@layer utilities {
  .tw\:text-red-500 {
    color: var(--tw-color-red-500);
  }
}
```

--------------------------------

### Use complete class names in HTML

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

To ensure Tailwind generates the correct CSS, always use full, static class names. This example shows the correct way to conditionally apply classes.

```html
<div class="{{ error ? 'text-red-600' : 'text-green-600' }}"></div>
```

--------------------------------

### Responsive minimum block size

Source: https://tailwindcss.com/docs/min-block-size

Prefix a `min-block-size` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="block-24 min-block-0 md:min-block-full ...">  <!-- ... --></div>
```

--------------------------------

### Responsive User Select

Source: https://tailwindcss.com/docs/user-select

Prefix an `user-select` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This allows for adaptive text selection behavior.

```html
<div class="select-none md:select-all ...">  <!-- ... --></div>
```

--------------------------------

### Apply Responsive Saturation Filters with Tailwind CSS

Source: https://tailwindcss.com/docs/filter-saturate

Apply saturation utilities conditionally at different screen sizes by prefixing them with breakpoint variants like `md:`, allowing for responsive design.

```html
<img class="saturate-50 md:saturate-150 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Setting Block Border Color (Fuchsia)

Source: https://tailwindcss.com/docs/border-color

Use the `border-block-fuchsia-{shade}` utility to set the color of the block borders. For example, `border-block-fuchsia-900` sets the block borders to a dark fuchsia.

```html
<div class="border-block-fuchsia-900">...</div>
```

--------------------------------

### Custom @supports Variant Configuration

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Configure custom shortcuts for common `@supports` rules by creating new variants in the `supports-*` namespace. This simplifies repetitive `@supports` usage.

```css
@custom-variant supports-grid {
  @supports (display: grid) {
    @slot;
  }
}
```

--------------------------------

### Setting Block Border Color (Purple)

Source: https://tailwindcss.com/docs/border-color

Use the `border-block-purple-{shade}` utility to set the color of the block borders. For example, `border-block-purple-900` sets the block borders to a dark purple.

```html
<div class="border-block-purple-900">...</div>
```

--------------------------------

### Overriding Component Classes with Utilities

Source: https://tailwindcss.com/docs/adding-custom-styles

Demonstrates how utility classes can override styles defined in custom component classes.

```html
<div class="card rounded-none">
  <!-- ... -->
</div>
```

--------------------------------

### Apply Basic Saturation Filters with Tailwind CSS

Source: https://tailwindcss.com/docs/filter-saturate

Use `saturate-<number>` utilities to control an element's saturation level, such as `saturate-50` for 50% saturation or `saturate-200` for 200% saturation.

```html
<img class="saturate-50 ..." src="/img/mountains.jpg" /><img class="saturate-100 ..." src="/img/mountains.jpg" /><img class="saturate-150 ..." src="/img/mountains.jpg" /><img class="saturate-200 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Basic Filter Application

Source: https://tailwindcss.com/docs/filter

Use utilities like `blur-xs` and `grayscale` to apply filters to an element. Multiple filter utilities can be combined.

```html
<img class="blur-xs" src="/img/mountains.jpg" /><img class="grayscale" src="/img/mountains.jpg" /><img class="blur-xs grayscale" src="/img/mountains.jpg" />
```

--------------------------------

### Setting Left Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-l-{color}-{shade}` utility to set the color of the left border. For example, `border-l-yellow-400` sets the left border to a medium yellow.

```html
<div class="border-l-yellow-400">...</div>
```

--------------------------------

### Setting Bottom Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-b-{color}-{shade}` utility to set the color of the bottom border. For example, `border-b-green-700` sets the bottom border to a dark green.

```html
<div class="border-b-green-700">...</div>
```

--------------------------------

### Applying Responsive Mask Mode (HTML)

Source: https://tailwindcss.com/docs/mask-mode

Use breakpoint variants like `md:` to apply `mask-mode` utilities only at specific screen sizes and above, enabling responsive design.

```html
<div class="mask-alpha md:mask-luminance ...">  <!-- ... --></div>
```

--------------------------------

### Setting Right Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-r-{color}-{shade}` utility to set the color of the right border. For example, `border-r-red-600` sets the right border to a dark red.

```html
<div class="border-r-red-600">...</div>
```

--------------------------------

### Responsive Position Utility

Source: https://tailwindcss.com/docs/position

Prefix a `position` utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<div class="relative md:absolute ...">
  <!-- ... -->
</div>
```

--------------------------------

### Setting Top Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-{color}-{shade}` utility to set the color of the top border. For example, `border-t-blue-500` sets the top border to a medium blue.

```html
<div class="border-t-blue-500">...</div>
```

--------------------------------

### Importing Google Fonts and configuring theme

Source: https://tailwindcss.com/docs/font-family

Import fonts from Google Fonts using @import and configure a corresponding theme variable. Ensure @import statements precede other rules.

```css
@import url("https://fonts.googleapis.com/css2?family=Roboto&display=swap");@import "tailwindcss";@theme {  --font-roboto: "Roboto", sans-serif; }
```

--------------------------------

### Configure PostCSS plugins

Source: https://tailwindcss.com/docs/installation/framework-guides/rspack/react

Create postcss.config.mjs and register the @tailwindcss/postcss plugin to enable Tailwind CSS processing.

```javascript
export default {  plugins: {    "@tailwindcss/postcss": {},  },};
```

--------------------------------

### Start-start logical border radius with custom property

Source: https://tailwindcss.com/docs/border-radius

Apply start-start border-radius using a custom CSS property variable.

```css
border-start-start-radius: var(<custom-property>);
```

--------------------------------

### Set Caret Color with Tailwind CSS

Source: https://tailwindcss.com/docs/caret-color

Use the `caret-{color}` utility class to set the color of the text cursor. This example shows how to use predefined mist colors.

```html
caret-mist-100
caret-color: var(--color-mist-100); /* oklch(96.3% 0.002 197.1) */
caret-mist-200
caret-color: var(--color-mist-200); /* oklch(92.5% 0.005 214.3) */
caret-mist-300
caret-color: var(--color-mist-300); /* oklch(87.2% 0.007 219.6) */
caret-mist-400
caret-color: var(--color-mist-400); /* oklch(72.3% 0.014 214.4) */
caret-mist-500
caret-color: var(--color-mist-500); /* oklch(56% 0.021 213.5) */
caret-mist-600
caret-color: var(--color-mist-600); /* oklch(45% 0.017 213.2) */
caret-mist-700
caret-color: var(--color-mist-700); /* oklch(37.8% 0.015 216) */
caret-mist-800
caret-color: var(--color-mist-800); /* oklch(27.5% 0.011 216.9) */
caret-mist-900
caret-color: var(--color-mist-900); /* oklch(21.8% 0.008 223.9) */
caret-mist-950
caret-color: var(--color-mist-950); /* oklch(14.8% 0.004 228.8) */
```

--------------------------------

### Responsive grid-template-rows with breakpoint variants

Source: https://tailwindcss.com/docs/grid-template-rows

Apply different row configurations at different screen sizes using breakpoint prefixes like md:. The grid uses 2 rows by default and 6 rows at medium screen sizes and above.

```html
<div class="grid grid-rows-2 md:grid-rows-6 ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive Cursor Utility in HTML

Source: https://tailwindcss.com/docs/cursor

Apply different cursor styles based on screen size using responsive variants like "md:cursor-auto".

```html
<button class="cursor-not-allowed md:cursor-auto ...">  <!-- ... --></button>
```

--------------------------------

### Registering a Basic Functional Utility

Source: https://tailwindcss.com/docs/adding-custom-styles

Register a simple functional utility that accepts an argument and uses the `--value()` function to resolve it.

```css
@utility tab-* {
  tab-size: --value(--tab-size-*);
}
```

--------------------------------

### Applying Filters Responsively

Source: https://tailwindcss.com/docs/filter

Use breakpoint variants like `md:` to apply filter utilities only at specific screen sizes and above. This allows for responsive filter designs.

```html
<img class="blur-sm md:filter-none ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive text-decoration-thickness

Source: https://tailwindcss.com/docs/text-decoration-thickness

Prefix a `text-decoration-thickness` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<p class="underline md:decoration-4 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Radial Gradient Mask From Custom Property

Source: https://tailwindcss.com/docs/mask-image

Applies a radial gradient mask, using a custom CSS property for the 'from' position. Enables dynamic control over the gradient's start.

```css
mask-image: radial-gradient(var(--tw-mask-radial-shape) var(--tw-mask-radial-size) at var(--tw-mask-radial-position), black var(<custom-property>), transparent var(--tw-mask-radial-to));
```

--------------------------------

### Use a custom blur utility class

Source: https://tailwindcss.com/docs/filter-blur

Apply a custom blur utility class, such as blur-2xs, after defining its value in the theme configuration.

```html
<img class="blur-2xs" src="/img/mountains.jpg" />
```

--------------------------------

### Basic min-width with spacing scale

Source: https://tailwindcss.com/docs/min-width

Apply fixed minimum widths using numeric utilities like min-w-24, min-w-40, etc., based on the spacing scale.

```html
<div class="w-20 ...">
  <div class="min-w-80 ...">min-w-80</div>
  <div class="min-w-64 ...">min-w-64</div>
  <div class="min-w-48 ...">min-w-48</div>
  <div class="min-w-40 ...">min-w-40</div>
  <div class="min-w-32 ...">min-w-32</div>
  <div class="min-w-24 ...">min-w-24</div>
</div>
```

--------------------------------

### Apply Responsive Overflow Utilities in HTML

Source: https://tailwindcss.com/docs/overflow

Prefix an `overflow` utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<div class="overflow-auto md:overflow-scroll ...">  <!-- ... --></div>
```

--------------------------------

### Start-start logical border radius reset and full

Source: https://tailwindcss.com/docs/border-radius

Remove start-start border-radius or apply infinite radius using logical properties.

```css
border-start-start-radius: 0;
```

```css
border-start-start-radius: calc(infinity * 1px);
```

--------------------------------

### Responsive Tab Size Design

Source: https://tailwindcss.com/docs/tab-size

Prefix a `tab-size` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<pre class="tab-4 md:tab-8 ...">  <!-- ... --></pre>
```

--------------------------------

### Basic fixed width with spacing scale

Source: https://tailwindcss.com/docs/width

Set fixed widths using w-<number> utilities like w-24 and w-96 based on the spacing scale.

```html
<div class="w-96 ...">w-96</div><div class="w-80 ...">w-80</div><div class="w-64 ...">w-64</div><div class="w-48 ...">w-48</div><div class="w-40 ...">w-40</div><div class="w-32 ...">w-32</div><div class="w-24 ...">w-24</div>
```

--------------------------------

### Setting Top Block Border Color

Source: https://tailwindcss.com/docs/border-color

Use the `border-t-{color}-{shade}` utility to set the color of the top border. For example, `border-t-blue-800` sets the top border to a dark blue.

```html
<div class="border-t-blue-800">...</div>
```

--------------------------------

### Apply responsive font sizes

Source: https://tailwindcss.com/docs/font-size

Use breakpoint prefixes like `md:` to apply font size utilities only at specific screen sizes and above. This enables responsive typography.

```html
<p class="text-sm md:text-base ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Custom min-width with bracket syntax

Source: https://tailwindcss.com/docs/min-width

Use the min-w-[<value>] syntax to set minimum width with arbitrary custom values.

```html
<div class="min-w-[220px] ...">
  <!-- ... -->
</div>
```

--------------------------------

### Applying Predefined Perspective Origin Utilities in HTML

Source: https://tailwindcss.com/docs/perspective-origin

Use utilities like `perspective-origin-top-left` and `perspective-origin-bottom-right` to control the vanishing point of a perspective on an element.

```html
<div class="size-20 perspective-near perspective-origin-top-left ...">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-x-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div><div class="size-20 perspective-near perspective-origin-bottom-right …">  <div class="translate-z-12 rotate-x-0 bg-sky-300/75 ...">1</div>  <div class="-translate-z-12 rotate-y-18 bg-sky-300/75 ...">2</div>  <div class="translate-x-12 rotate-y-90 bg-sky-300/75 ...">3</div>  <div class="-translate-z-12 -rotate-y-90 bg-sky-300/75 ...">4</div>  <div class="-translate-y-12 rotate-x-90 bg-sky-300/75 ...">5</div>  <div class="translate-y-12 -rotate-x-90 bg-sky-300/75 ...">6</div></div>
```

--------------------------------

### Styling at Breakpoints

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Prefix utility classes with breakpoint names (e.g., 'sm:') to apply styles only at specific screen sizes and above.

```html
<div class="grid grid-cols-2 sm:grid-cols-3">
  <!-- ... -->
</div>
```

--------------------------------

### Safelisting utilities with ranges

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Leverage brace expansion in `@source inline()` to generate multiple classes at once, including ranges and variants.

```css
@import "tailwindcss";@source inline("{hover:,}bg-red-{50,{100..900..100},950}");
```

--------------------------------

### Apply responsive blur filters

Source: https://tailwindcss.com/docs/filter-blur

Use responsive prefixes like md: to apply blur utilities conditionally based on screen size, such as md:blur-lg.

```html
<img class="blur-none md:blur-lg ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive Text Wrap with Breakpoints

Source: https://tailwindcss.com/docs/text-wrap

Prefix a text-wrap utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<h1 class="text-pretty md:text-balance ...">  </h1>
```

--------------------------------

### Apply responsive object-fit with breakpoint variants (HTML)

Source: https://tailwindcss.com/docs/object-fit

Prefix an `object-fit` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<img class="object-contain md:object-cover" src="/img/mountains.jpg" />
```

--------------------------------

### Set Fixed Height with Tailwind CSS

Source: https://tailwindcss.com/docs/height

Use `h-<number>` utilities to set an element to a fixed height based on the Tailwind spacing scale. Examples include `h-24` to `h-96`.

```html
<div class="h-96 ...">h-96</div><div class="h-80 ...">h-80</div><div class="h-64 ...">h-64</div><div class="h-48 ...">h-48</div><div class="h-40 ...">h-40</div><div class="h-32 ...">h-32</div><div class="h-24 ...">h-24</div>
```

--------------------------------

### Set Border Width Using Logical Properties (Start/End)

Source: https://tailwindcss.com/docs/border-width

Utilize `border-s-*` and `border-e-*` for `border-inline-start-width` and `border-inline-end-width`. These properties automatically adjust to the text direction (LTR or RTL).

```html
<div dir="ltr">  <div class="border-s-4 ..."></div></div><div dir="rtl">  <div class="border-s-4 ..."></div></div>
```

--------------------------------

### Custom Data Attribute Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Define custom variants for common data attributes to simplify their usage. This example creates a `data-checked` variant for a `data-ui~="checked"` attribute.

```css
@import "tailwindcss";@custom-variant data-checked (&[data-ui~="checked"]);
```

--------------------------------

### Applying Responsive mask-clip Utilities in HTML

Source: https://tailwindcss.com/docs/mask-clip

Apply `mask-clip` utilities conditionally at different screen sizes using responsive prefixes like `md:`.

```html
<div class="mask-clip-border md:mask-clip-padding ...">  <!-- ... --></div>
```

--------------------------------

### Restoring forced colors

Source: https://tailwindcss.com/docs/forced-color-adjust

Use the `forced-color-adjust-auto` utility to make an element adhere to colors enforced by forced colors mode. This can be useful to undo `forced-color-adjust-none`, for example on a larger screen size.

```html
<form>
  <fieldset class="forced-color-adjust-none lg:forced-color-adjust-auto ...">
    <legend>Choose a color:</legend>
    <select class="hidden lg:block">
      <option value="White">White</option>
      <option value="Gray">Gray</option>
      <option value="Black">Black</option>
    </select>
    <div class="lg:hidden">
      <label>
        <input class="sr-only" type="radio" name="color-choice" value="White" />
        <!-- ... -->
      </label>
      <!-- ... -->
    </div>
  </fieldset>
</form>
```

--------------------------------

### Responsive backdrop brightness with breakpoint variant

Source: https://tailwindcss.com/docs/backdrop-filter-brightness

Prefix backdrop-filter brightness utilities with breakpoint variants like md: to apply the utility only at specific screen sizes and above.

```html
<div class="backdrop-brightness-110 md:backdrop-brightness-150 ...">  <!-- ... --></div>
```

--------------------------------

### Customize Column Theme Variables

Source: https://tailwindcss.com/docs/columns

Modify fixed-width column utilities by updating `--container-*` theme variables in your Tailwind CSS configuration. This example sets `--container-4xs` to `14rem`.

```css
@theme {
  --container-4xs: 14rem;
}
```

--------------------------------

### Migrate Space-between to Flexbox Gap

Source: https://tailwindcss.com/docs/upgrade-guide

Consider migrating `space-x-*` and `space-y-*` utilities to flexbox with `gap` for improved performance on large pages.

```html
<div class="space-y-4 p-4"><div class="flex flex-col gap-4 p-4">
  <label for="name">Name</label>
  <input type="text" name="name" /></div>
```

--------------------------------

### Basic background-clip utilities

Source: https://tailwindcss.com/docs/background-clip

Use `bg-clip-border`, `bg-clip-padding`, and `bg-clip-content` to control the bounding box of an element's background.

```html
<div class="border-4 bg-indigo-500 bg-clip-border p-3"></div><div class="border-4 bg-indigo-500 bg-clip-padding p-3"></div><div class="border-4 bg-indigo-500 bg-clip-content p-3"></div>
```

--------------------------------

### Apply Left Mask Gradient From Custom Property

Source: https://tailwindcss.com/docs/mask-image

Use `mask-l-from-(<custom-property>)` to apply a left mask gradient. The gradient starts from black at a position defined by a custom property and transitions to transparent.

```css
mask-image: linear-gradient(to left, black var(<custom-property>), transparent var(--tw-mask-left-to));
```

--------------------------------

### Apply Left Mask Gradient From Number

Source: https://tailwindcss.com/docs/mask-image

Use `mask-l-from-<number>` to apply a left mask gradient. The gradient starts from black at a calculated position based on `--spacing` and transitions to transparent.

```css
mask-image: linear-gradient(to left, black calc(var(--spacing) * <number>), transparent var(--tw-mask-left-to));
```

--------------------------------

### Predefined inset shadows

Source: https://tailwindcss.com/docs/box-shadow

Apply predefined inset box-shadows using utility classes like `inset-shadow-2xs`, `inset-shadow-xs`, and `inset-shadow-sm`. The `inset-shadow-none` class removes any existing inset shadow.

```html
<div class="inset-shadow-2xs">...</div>
<div class="inset-shadow-xs">...</div>
<div class="inset-shadow-sm">...</div>
<div class="inset-shadow-none">...</div>
```

--------------------------------

### Apply Bottom Mask Gradient from Custom Property

Source: https://tailwindcss.com/docs/mask-image

Use `mask-b-from-(<custom-property>)` to apply a bottom mask gradient. The gradient starts from black at a position defined by a custom property and transitions to transparent.

```css
mask-image: linear-gradient(to bottom, black var(<custom-property>), transparent var(--tw-mask-bottom-to));
```

--------------------------------

### Apply Responsive Border Width

Source: https://tailwindcss.com/docs/border-width

Prefix border width utilities with responsive variants like `md:` to apply them only at specific screen sizes and above. This enables adaptive designs.

```html
<div class="border-2 md:border-t-4 ...">  <!-- ... --></div>
```

--------------------------------

### Apply Bottom Mask Gradient from Number

Source: https://tailwindcss.com/docs/mask-image

Use `mask-b-from-<number>` to apply a bottom mask gradient. The gradient starts from black at a calculated position based on `--spacing` and transitions to transparent.

```css
mask-image: linear-gradient(to bottom, black calc(var(--spacing) * <number>), transparent var(--tw-mask-bottom-to));
```

--------------------------------

### Logical Border Properties

Source: https://tailwindcss.com/docs/border-color

Use utilities like `border-s-indigo-500` and `border-e-indigo-500` for logical start/end properties, and `border-bs-indigo-500` and `border-be-indigo-500` for block start/end properties.

```html
<div dir="ltr">  <div class="border-s-indigo-500 ..."></div></div><div dir="rtl">  <div class="border-s-indigo-500 ..."></div></div>
```

```html
<div class="border-bs-indigo-500 ..."></div>
```

--------------------------------

### Responsive Aspect Ratio

Source: https://tailwindcss.com/docs/aspect-ratio

Prefix an `aspect-ratio` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<iframe class="aspect-video md:aspect-square ..." src="https://www.youtube.com/embed/dQw4w9WgXcQ"></iframe>
```

--------------------------------

### Setting border-bottom color with predefined colors

Source: https://tailwindcss.com/docs/border-color

Use these utilities to set the border-bottom color using Tailwind's predefined color palette. Examples include olive, mist, and taupe shades.

```html
<div class="border-b-olive-300"></div>
<div class="border-b-mist-50"></div>
<div class="border-b-taupe-100"></div>
```

--------------------------------

### Using logical float properties

Source: https://tailwindcss.com/docs/float

Employ `float-start` and `float-end` for direction-agnostic floating. These utilities adapt to the text direction, mapping to left or right as appropriate.

```html
<article>  <img class="float-start ..." src="/img/mountains.jpg" />  <p>Maybe we can live without libraries, people like you and me. ...</p></article><article dir="rtl">  <img class="float-start ..." src="/img/mountains.jpg" />  <p>... ربما يمكننا العيش بدون مكتبات، أشخاص مثلي ومثلك. ربما. بالتأكيد</p></article>
```

--------------------------------

### Apply visibility utilities at specific breakpoints

Source: https://tailwindcss.com/docs/visibility

Prefix a `visibility` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="visible md:invisible ...">  <!-- ... --></div>
```

--------------------------------

### Import Tailwind CSS into app.css

Source: https://tailwindcss.com/docs/installation/framework-guides/adonisjs

Import Tailwind CSS styles into your main `app.css` file and configure it to scan your views directory for utility classes.

```CSS
@import "tailwindcss";@source "../views";
```

--------------------------------

### Responsive Backdrop Blur

Source: https://tailwindcss.com/docs/backdrop-filter-blur

Apply backdrop blur utilities responsively by prefixing them with a breakpoint variant, such as `md:`. This ensures the blur effect is applied only at medium screen sizes and above.

```html
<div class="backdrop-blur-none md:backdrop-blur-lg ...">  <!-- ... --></div>
```

--------------------------------

### Applying responsive flex-grow utilities in Tailwind CSS

Source: https://tailwindcss.com/docs/flex-grow

Use breakpoint prefixes like `md:` to apply `flex-grow` utilities conditionally, allowing items to grow or not grow based on screen size.

```html
<div class="grow md:grow-0 ...">  <!-- ... --></div>
```

--------------------------------

### Changing Gradient Interpolation Mode

Source: https://tailwindcss.com/docs/background-image

Use the interpolation modifier to control how colors in a gradient are interpolated. Examples show srgb, hsl, oklab, oklch, longer, shorter, increasing, and decreasing modes.

```html
<div class="bg-linear-to-r/srgb from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/hsl from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/oklab from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/oklch from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/longer from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/shorter from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/increasing from-indigo-500 to-teal-400"></div><div class="bg-linear-to-r/decreasing from-indigo-500 to-teal-400"></div>
```

--------------------------------

### Applying text decoration with responsive design

Source: https://tailwindcss.com/docs/text-decoration-line

Prefix a `text-decoration-line` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<a class="no-underline md:underline ..." href="...">  <!-- ... --></a>
```

--------------------------------

### Basic Cursor Utilities in HTML

Source: https://tailwindcss.com/docs/cursor

Apply standard cursor styles like "cursor-pointer", "cursor-progress", and "cursor-not-allowed" to HTML elements to indicate interaction states.

```html
<button class="cursor-pointer ...">Submit</button><button class="cursor-progress ...">Saving...</button><button class="cursor-not-allowed ..." disabled>Confirm</button>
```

--------------------------------

### Using CSS Variables Instead of theme() Function

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, CSS variables are recommended over the `theme()` function for accessing theme values. This example shows direct usage of a CSS variable for background color.

```css
.my-class {
  background-color: theme(colors.red.500);
  background-color: var(--color-red-500);
}
```

--------------------------------

### Responsive pseudo-element content

Source: https://tailwindcss.com/docs/content

Apply `content` utilities at specific breakpoints using prefix variants like `md:`. This allows for different content to be displayed on different screen sizes.

```html
<p class="before:content-['Mobile'] md:before:content-['Desktop'] ..."></p>
```

--------------------------------

### Responsive list-style-image

Source: https://tailwindcss.com/docs/list-style-image

Prefix a `list-style-image` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<ul class="list-image-none md:list-image-[url(/img/checkmark.png)] ...">  <!-- ... --></ul>
```

--------------------------------

### Responsive Invert Design

Source: https://tailwindcss.com/docs/filter-invert

Control color inversion at different screen sizes by prefixing utilities with breakpoint variants like `md:`.

```html
<img class="invert md:invert-0 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive Flex Wrap Design

Source: https://tailwindcss.com/docs/flex-wrap

Prefix a `flex-wrap` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="flex flex-wrap md:flex-wrap-reverse ...">  <!-- ... --></div>
```

--------------------------------

### Attempting to style based on a subsequent peer (won't work)

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Demonstrates that the `peer` marker only works on previous siblings due to CSS's subsequent-sibling combinator limitations. Styling based on a subsequent peer will not function as expected.

```html
<label>
  <span class="peer-invalid:text-red-500 ...">Email</span>
  <input type="email" class="peer ..." /></label>
```

--------------------------------

### Apply Responsive Maximum Width with Tailwind CSS Breakpoints

Source: https://tailwindcss.com/docs/max-width

Prefix `max-width` utilities with breakpoint variants like `md:` to apply them only at specific screen sizes and above.

```html
<div class="max-w-sm md:max-w-lg ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Backdrop Contrast Filter

Source: https://tailwindcss.com/docs/backdrop-filter-contrast

Prefix a `backdrop-filter: contrast()` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="backdrop-contrast-125 md:backdrop-contrast-150 ...">  <!-- ... --></div>
```

--------------------------------

### Custom grid-auto-columns Value

Source: https://tailwindcss.com/docs/grid-auto-columns

Shows how to set implicit grid column sizes using a custom value with the `auto-cols-[<value>]` syntax. This is useful for precise control or when using values not directly supported by utility classes.

```html
<div class="auto-cols-[minmax(0,2fr)] ...">  <!-- ... --></div>
```

--------------------------------

### Flex Container Layout

Source: https://tailwindcss.com/docs/display

Use the 'flex' utility to establish a block-level flex container for arranging child elements.

```html
<div class="flex items-center">  <img src="path/to/image.jpg" />  <div>    <strong>Andrew Alfred</strong>    <span>Technical advisor</span>  </div></div>
```

--------------------------------

### Parent/Sibling Styling with ARIA Variants

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Utilize `group-aria-*` and `peer-aria-*` variants to style parent or sibling elements based on ARIA states. This example rotates an SVG arrow in a table header based on `aria-sort`.

```html
<table>  <thead>    <tr>    <th aria-sort="ascending" class="group">      Invoice #      <svg class="group-aria-[sort=ascending]:rotate-0 group-aria-[sort=descending]:rotate-180"><!-- ... --></svg>    </th>    <!-- ... -->    </tr>  </thead>  <!-- ... --></table>
```

--------------------------------

### Responsive Grid Columns

Source: https://tailwindcss.com/docs/grid-template-columns

Apply grid-template-columns utilities responsively by prefixing with a breakpoint variant like 'md:'. This ensures the grid layout adapts to different screen sizes.

```html
<div class="grid grid-cols-1 md:grid-cols-6 ...">  <!-- ... --></div>
```

--------------------------------

### Get Resolved CSS Variable Value in JavaScript

Source: https://tailwindcss.com/docs/theme

Use getComputedStyle() to retrieve the resolved value of a theme variable from the document root. Necessary when you need the actual computed value rather than the variable reference.

```JavaScript
let styles = getComputedStyle(document.documentElement);
let shadow = styles.getPropertyValue("--shadow-xl");
```

--------------------------------

### Apply Vertical Mask Gradient From Color

Source: https://tailwindcss.com/docs/mask-image

Use `mask-y-from-<color>` to apply vertical mask gradients (top and bottom). The gradients start from the specified color and transition to transparent. `mask-composite: intersect;` is used to combine them.

```css
mask-image: linear-gradient(to top, <color> var(--tw-mask-top-from), transparent var(--tw-mask-top-to)), linear-gradient(to bottom, <color> var(--tw-mask-bottom-from), transparent var(--tw-mask-bottom-to)); mask-composite: intersect;
```

--------------------------------

### Configure Source Detection for Utilities

Source: https://tailwindcss.com/docs/preflight

When importing Tailwind CSS files individually, add `source(none)` to the `utilities.css` import to configure source detection, which affects generated utilities.

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/utilities.css" layer(utilities);
@import "tailwindcss/utilities.css" layer(utilities) source(none);
```

--------------------------------

### Inset shadow color utilities with predefined colors

Source: https://tailwindcss.com/docs/box-shadow

Apply predefined color values to inset shadows using classes like `inset-shadow-black`, `inset-shadow-white`, and various shades of red, orange, amber, and yellow.

```html
<div class="inset-shadow-black">...</div>
<div class="inset-shadow-white">...</div>
<div class="inset-shadow-red-50">...</div>
<div class="inset-shadow-red-100">...</div>
<div class="inset-shadow-red-200">...</div>
<div class="inset-shadow-red-300">...</div>
<div class="inset-shadow-red-400">...</div>
<div class="inset-shadow-red-500">...</div>
<div class="inset-shadow-red-600">...</div>
<div class="inset-shadow-red-700">...</div>
<div class="inset-shadow-red-800">...</div>
<div class="inset-shadow-red-900">...</div>
<div class="inset-shadow-red-950">...</div>
<div class="inset-shadow-orange-50">...</div>
<div class="inset-shadow-orange-100">...</div>
<div class="inset-shadow-orange-200">...</div>
<div class="inset-shadow-orange-300">...</div>
<div class="inset-shadow-orange-400">...</div>
<div class="inset-shadow-orange-500">...</div>
<div class="inset-shadow-orange-600">...</div>
<div class="inset-shadow-orange-700">...</div>
<div class="inset-shadow-orange-800">...</div>
<div class="inset-shadow-orange-900">...</div>
<div class="inset-shadow-orange-950">...</div>
<div class="inset-shadow-amber-50">...</div>
<div class="inset-shadow-amber-100">...</div>
<div class="inset-shadow-amber-200">...</div>
<div class="inset-shadow-amber-300">...</div>
<div class="inset-shadow-amber-400">...</div>
<div class="inset-shadow-amber-500">...</div>
<div class="inset-shadow-amber-600">...</div>
<div class="inset-shadow-amber-700">...</div>
<div class="inset-shadow-amber-800">...</div>
<div class="inset-shadow-amber-900">...</div>
<div class="inset-shadow-amber-950">...</div>
<div class="inset-shadow-yellow-50">...</div>
<div class="inset-shadow-yellow-100">...</div>
<div class="inset-shadow-yellow-200">...</div>
<div class="inset-shadow-yellow-300">...</div>
```

--------------------------------

### Growing flex items proportionally with Tailwind CSS `grow-<number>`

Source: https://tailwindcss.com/docs/flex-grow

Apply `grow-<number>` utilities, such as `grow-3` or `grow-7`, to make flex items grow relative to each other based on their specified growth factor.

```html
<div class="flex ...">  <div class="size-14 grow-3 ...">01</div>  <div class="size-14 grow-7 ...">02</div>  <div class="size-14 grow-3 ...">03</div></div>
```

--------------------------------

### Extending Tailwind CSS Theme with Custom Font

Source: https://tailwindcss.com/docs/theme

This example demonstrates how to extend the default Tailwind CSS theme by defining a new custom font variable, `--font-script`, using the `@theme` directive in `app.css`.

```css
@import "tailwindcss";@theme {  --font-script: Great Vibes, cursive;}
```

--------------------------------

### Applying scrollbar width responsively

Source: https://tailwindcss.com/docs/scrollbar-width

Prefix a `scrollbar-width` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="scrollbar-none md:scrollbar-auto ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Drop Shadows

Source: https://tailwindcss.com/docs/filter-drop-shadow

Apply drop shadow utilities at specific breakpoints by prefixing them with a breakpoint variant like `md:`.

```html
<svg class="drop-shadow-md md:drop-shadow-xl ...">  <!-- ... --></svg>
```

--------------------------------

### Arbitrary peer variants with custom selectors

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Create one-off `peer-*` variants by providing a custom selector in square brackets, like `peer-[.selector]`. This allows for dynamic styling based on arbitrary sibling states.

```html
<form>
  <label for="email">Email:</label>
  <input id="email" name="email" type="email" class="is-dirty peer" required />
  <div class="peer-[.is-dirty]:peer-required:block hidden">This field is required.</div>
  <!-- ... -->
</form>
```

--------------------------------

### Apply Vertical Mask Gradient From Value

Source: https://tailwindcss.com/docs/mask-image

Use `mask-y-from-[<value>]` to apply vertical mask gradients (top and bottom). The gradients start from black at specified values and transition to transparent. `mask-composite: intersect;` is used to combine them.

```css
mask-image: linear-gradient(to top, black <value>, transparent var(--tw-mask-top-to)), linear-gradient(to bottom, black <value>, transparent var(--tw-mask-bottom-to)); mask-composite: intersect;
```

--------------------------------

### Percentage-Based Inline Sizes

Source: https://tailwindcss.com/docs/inline-size

Use `inline-full` or `inline-<fraction>` utilities to give an element a percentage-based inline size.

```html
<div class="flex ...">  <div class="inline-1/2 ...">inline-1/2</div>  <div class="inline-1/2 ...">inline-1/2</div></div><div class="flex ...">  <div class="inline-2/5 ...">inline-2/5</div>  <div class="inline-3/5 ...">inline-3/5</div></div><div class="flex ...">  <div class="inline-1/3 ...">inline-1/3</div>  <div class="inline-2/3 ...">inline-2/3</div></div><div class="inline-full ...">inline-full</div>
```

--------------------------------

### Apply Vertical Mask Gradient From Percentage

Source: https://tailwindcss.com/docs/mask-image

Use `mask-y-from-<percentage>` to apply vertical mask gradients (top and bottom). The gradients start from black at the specified percentage and transition to transparent. `mask-composite: intersect;` is used to combine them.

```css
mask-image: linear-gradient(to top, black <percentage>, transparent var(--tw-mask-top-to)), linear-gradient(to bottom, black <percentage>, transparent var(--tw-mask-bottom-to)); mask-composite: intersect;
```

--------------------------------

### Responsive align-self utility

Source: https://tailwindcss.com/docs/align-self

Prefix an `align-self` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="self-auto md:self-end ...">  <!-- ... --></div>
```

--------------------------------

### Set Right Border Color with Predefined Colors

Source: https://tailwindcss.com/docs/border-color

Use `border-r-{color}-{shade}` utilities to set the right border color using Tailwind's predefined color palette. Examples include `border-r-red-500` and `border-r-orange-100`.

```html
border-r-red-50
border-r-red-100
border-r-red-200
border-r-red-300
border-r-red-400
border-r-red-500
border-r-red-600
border-r-red-700
border-r-red-800
border-r-red-900
border-r-red-950
border-r-orange-50
border-r-orange-100
```

--------------------------------

### Stacking Variants for Specific States

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Demonstrates stacking multiple variants like `dark:`, `md:`, and `hover:` to apply styles under specific conditions, such as dark mode, medium screen size, and hover.

```html
<button class="dark:md:hover:bg-fuchsia-600 ...">Save changes</button>
```

--------------------------------

### Using Logical Properties for Text Alignment

Source: https://tailwindcss.com/docs/text-align

Use `text-start` and `text-end` utilities for text alignment that respects the text direction (LTR/RTL). These map to left/right alignment based on the `dir` attribute.

```html
<div dir="rtl" lang="ar">  <p class="text-end">فبدأت بالسير نحو الماء...</p></div>
```

--------------------------------

### Use Auto Select Behavior

Source: https://tailwindcss.com/docs/user-select

Use the `select-auto` utility to revert to the default browser behavior for text selection. This is useful when you want the browser to handle selection logic.

```html
<div class="select-auto ...">The quick brown fox jumps over the lazy dog.</div>
```

--------------------------------

### Override Image Constraints

Source: https://tailwindcss.com/docs/preflight

Use the `max-w-none` utility to override the default responsive constraints for images and videos.

```html
<img class="max-w-none" src="..." alt="..." />
```

--------------------------------

### Setting Inline Border Color to Lime Shades

Source: https://tailwindcss.com/docs/border-color

Shows how to apply inline border colors using various shades of lime, from `lime-50` to `lime-950`.

```html
<!-- Set border-inline-color to lime-50 -->
<div class="border-x-lime-50"></div>

<!-- Set border-inline-color to lime-100 -->
<div class="border-x-lime-100"></div>

<!-- Set border-inline-color to lime-200 -->
<div class="border-x-lime-200"></div>

<!-- Set border-inline-color to lime-300 -->
<div class="border-x-lime-300"></div>

<!-- Set border-inline-color to lime-400 -->
<div class="border-x-lime-400"></div>

<!-- Set border-inline-color to lime-500 -->
<div class="border-x-lime-500"></div>

<!-- Set border-inline-color to lime-600 -->
<div class="border-x-lime-600"></div>

<!-- Set border-inline-color to lime-700 -->
<div class="border-x-lime-700"></div>

<!-- Set border-inline-color to lime-800 -->
<div class="border-x-lime-800"></div>

<!-- Set border-inline-color to lime-900 -->
<div class="border-x-lime-900"></div>

<!-- Set border-inline-color to lime-950 -->
<div class="border-x-lime-950"></div>
```

--------------------------------

### Setting Inline Border Color with Tailwind CSS

Source: https://tailwindcss.com/docs/border-color

Use the `border-inline-color` utility to set the color of the inline borders (left and right in LTR, right and left in RTL). This example shows setting the border to a specific shade of orange.

```html
<!-- Set border-inline-color to orange-950 -->
<div class="border-x-orange-950"></div>
```

--------------------------------

### Dynamic ARIA Variants with Arbitrary Values

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use arbitrary values within square brackets to generate dynamic ARIA variants for complex or one-off ARIA attributes. This example styles table headers based on `aria-sort` values.

```html
<table>  <thead>    <tr>      <th        aria-sort="ascending"        class="aria-[sort=ascending]:bg-[url('/img/down-arrow.svg')] aria-[sort=descending]:bg-[url('/img/up-arrow.svg')]"      >        Invoice #      </th>      <!-- ... -->    </tr>  </thead>  <!-- ... --></table>
```

--------------------------------

### Apply responsive maximum height

Source: https://tailwindcss.com/docs/max-height

Use responsive prefixes like md: to apply max-height utilities conditionally based on screen size.

```html
<div class="h-48 max-h-full md:max-h-screen ...">  <!-- ... --></div>
```

--------------------------------

### Styling Based on Descendant State with group-has Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `group-has-*` variant to style an element based on the descendants of a parent element marked with the `group` class. This example shows an SVG that is hidden unless a descendant link is present.

```html
<div class="group ...">  <img src="..." />  <h4>Spencer Sharp</h4>  <svg class="hidden group-has-[a]:block ..."><!-- ... --></svg>  <p>Product Designer at <a href="...">planeteria.tech</a></p></div>
```

--------------------------------

### Set Top Border Color with Predefined Colors

Source: https://tailwindcss.com/docs/border-color

Use `border-t-{color}-{shade}` utilities to set the top border color using Tailwind's predefined color palette. For example, `border-t-olive-500` applies a specific shade of olive.

```html
border-t-olive-50
border-t-olive-100
border-t-olive-200
border-t-olive-300
border-t-olive-400
border-t-olive-500
border-t-olive-600
border-t-olive-700
border-t-olive-800
border-t-olive-900
border-t-olive-950
border-t-mist-50
border-t-mist-100
border-t-mist-200
border-t-mist-300
border-t-mist-400
border-t-mist-500
border-t-mist-600
border-t-mist-700
border-t-mist-800
border-t-mist-900
border-t-mist-950
border-t-taupe-50
border-t-taupe-100
border-t-taupe-200
border-t-taupe-300
border-t-taupe-400
border-t-taupe-500
border-t-taupe-600
border-t-taupe-700
border-t-taupe-800
border-t-taupe-900
border-t-taupe-950
```

--------------------------------

### Basic z-index

Source: https://tailwindcss.com/docs/z-index

Use z-<number> utilities like z-10 and z-50 to control the stack order of elements.

```html
<div class="z-40 ...">05</div><div class="z-30 ...">04</div><div class="z-20 ...">03</div><div class="z-10 ...">02</div><div class="z-0 ...">01</div>
```

--------------------------------

### max-inline-size Utilities using Container Scale

Source: https://tailwindcss.com/docs/max-inline-size

Set a fixed maximum inline size using utilities like `max-inline-sm` and `max-inline-xl` that are based on the container scale. This allows for consistent sizing across different container breakpoints.

```html
<div class="max-inline-md ...">  <!-- ... --></div>
```

--------------------------------

### Apply Vertical Mask Gradient From Custom Property

Source: https://tailwindcss.com/docs/mask-image

Use `mask-y-from-(<custom-property>)` to apply vertical mask gradients (top and bottom). The gradients start from black at positions defined by custom properties and transition to transparent. `mask-composite: intersect;` is used to combine them.

```css
mask-image: linear-gradient(to top, black var(<custom-property>), transparent var(--tw-mask-top-to)), linear-gradient(to bottom, black var(<custom-property>), transparent var(--tw-mask-bottom-to)); mask-composite: intersect;
```

--------------------------------

### Apply Vertical Mask Gradient From Number

Source: https://tailwindcss.com/docs/mask-image

Use `mask-y-from-<number>` to apply vertical mask gradients (top and bottom). The gradients start from black at calculated positions based on `--spacing` and transition to transparent. `mask-composite: intersect;` is used to combine them.

```css
mask-image: linear-gradient(to top, black calc(var(--spacing) * <number>), transparent var(--tw-mask-top-to)), linear-gradient(to bottom, black calc(var(--spacing) * <number>), transparent var(--tw-mask-bottom-to)); mask-composite: intersect;
```

--------------------------------

### CSS Custom Properties for Default Color Palette

Source: https://tailwindcss.com/docs/colors

These CSS custom properties define the default color palette, using `oklch` and hex values. They can be reused or remapped, for example, to redefine a color scale like `--color-gray-*` using a different base like `--color-slate-*`.

```css
--color-zinc-50: oklch(98.5% 0 0);  --color-zinc-100: oklch(96.7% 0.001 286.375);  --color-zinc-200: oklch(92% 0.004 286.32);  --color-zinc-300: oklch(87.1% 0.006 286.286);  --color-zinc-400: oklch(70.5% 0.015 286.067);  --color-zinc-500: oklch(55.2% 0.016 285.938);  --color-zinc-600: oklch(44.2% 0.017 285.786);  --color-zinc-700: oklch(37% 0.013 285.805);  --color-zinc-800: oklch(27.4% 0.006 286.033);  --color-zinc-900: oklch(21% 0.006 285.885);  --color-zinc-950: oklch(14.1% 0.005 285.823);  --color-neutral-50: oklch(98.5% 0 0);  --color-neutral-100: oklch(97% 0 0);  --color-neutral-200: oklch(92.2% 0 0);  --color-neutral-300: oklch(87% 0 0);  --color-neutral-400: oklch(70.8% 0 0);  --color-neutral-500: oklch(55.6% 0 0);  --color-neutral-600: oklch(43.9% 0 0);  --color-neutral-700: oklch(37.1% 0 0);  --color-neutral-800: oklch(26.9% 0 0);  --color-neutral-900: oklch(20.5% 0 0);  --color-neutral-950: oklch(14.5% 0 0);  --color-stone-50: oklch(98.5% 0.001 106.423);  --color-stone-100: oklch(97% 0.001 106.424);  --color-stone-200: oklch(92.3% 0.003 48.717);  --color-stone-300: oklch(86.9% 0.005 56.366);  --color-stone-400: oklch(70.9% 0.01 56.259);  --color-stone-500: oklch(55.3% 0.013 58.071);  --color-stone-600: oklch(44.4% 0.011 73.639);  --color-stone-700: oklch(37.4% 0.01 67.558);  --color-stone-800: oklch(26.8% 0.007 34.298);  --color-stone-900: oklch(21.6% 0.006 56.043);  --color-stone-950: oklch(14.7% 0.004 49.25);  --color-mauve-50: oklch(98.5% 0 0);  --color-mauve-100: oklch(96% 0.003 325.6);  --color-mauve-200: oklch(92.2% 0.005 325.62);  --color-mauve-300: oklch(86.5% 0.012 325.68);  --color-mauve-400: oklch(71.1% 0.019 323.02);  --color-mauve-500: oklch(54.2% 0.034 322.5);  --color-mauve-600: oklch(43.5% 0.029 321.78);  --color-mauve-700: oklch(36.4% 0.029 323.89);  --color-mauve-800: oklch(26.3% 0.024 320.12);  --color-mauve-900: oklch(21.2% 0.019 322.12);  --color-mauve-950: oklch(14.5% 0.008 326);  --color-olive-50: oklch(98.8% 0.003 106.5);  --color-olive-100: oklch(96.6% 0.005 106.5);  --color-olive-200: oklch(93% 0.007 106.5);  --color-olive-300: oklch(88% 0.011 106.6);  --color-olive-400: oklch(73.7% 0.021 106.9);  --color-olive-500: oklch(58% 0.031 107.3);  --color-olive-600: oklch(46.6% 0.025 107.3);  --color-olive-700: oklch(39.4% 0.023 107.4);  --color-olive-800: oklch(28.6% 0.016 107.4);  --color-olive-900: oklch(22.8% 0.013 107.4);  --color-olive-950: oklch(15.3% 0.006 107.1);  --color-mist-50: oklch(98.7% 0.002 197.1);  --color-mist-100: oklch(96.3% 0.002 197.1);  --color-mist-200: oklch(92.5% 0.005 214.3);  --color-mist-300: oklch(87.2% 0.007 219.6);  --color-mist-400: oklch(72.3% 0.014 214.4);  --color-mist-500: oklch(56% 0.021 213.5);  --color-mist-600: oklch(45% 0.017 213.2);  --color-mist-700: oklch(37.8% 0.015 216);  --color-mist-800: oklch(27.5% 0.011 216.9);  --color-mist-900: oklch(21.8% 0.008 223.9);  --color-mist-950: oklch(14.8% 0.004 228.8);  --color-taupe-50: oklch(98.6% 0.002 67.8);  --color-taupe-100: oklch(96% 0.002 17.2);  --color-taupe-200: oklch(92.2% 0.005 34.3);  --color-taupe-300: oklch(86.8% 0.007 39.5);  --color-taupe-400: oklch(71.4% 0.014 41.2);  --color-taupe-500: oklch(54.7% 0.021 43.1);  --color-taupe-600: oklch(43.8% 0.017 39.3);  --color-taupe-700: oklch(36.7% 0.016 35.7);  --color-taupe-800: oklch(26.8% 0.011 36.5);  --color-taupe-900: oklch(21.4% 0.009 43.1);  --color-taupe-950: oklch(14.7% 0.004 49.3);  --color-black: #000;  --color-white: #fff;
```

--------------------------------

### Responsive border-spacing

Source: https://tailwindcss.com/docs/border-spacing

Apply border spacing utilities at specific breakpoints using responsive prefixes like `md:`. This allows the spacing to adapt to different screen sizes.

```html
<table class="border-spacing-2 md:border-spacing-4 ...">  <!-- ... --></table>
```

--------------------------------

### Responsive Background Image

Source: https://tailwindcss.com/docs/background-image

Apply background image utilities at specific breakpoints by prefixing them with a breakpoint variant like `md:`. This ensures the utility only applies at medium screen sizes and above.

```html
<div class="from-purple-400 md:from-yellow-500 ...">  <!-- ... --></div>
```

--------------------------------

### Apply responsive animation with Tailwind CSS

Source: https://tailwindcss.com/docs/animation

Prefix animation utilities with breakpoint variants like `md:` to apply animations only at specific screen sizes.

```html
<div class="animate-none md:animate-spin ...">  <!-- ... --></div>
```

--------------------------------

### Using custom grid-template-rows values

Source: https://tailwindcss.com/docs/grid-template-rows

Apply completely custom row track values using the grid-rows-[<value>] bracket syntax for pixel-based or complex minmax definitions.

```html
<div class="grid-rows-[200px_minmax(900px,1fr)_100px] ...">
  <!-- ... -->
</div>
```

--------------------------------

### Configure PostCSS for Tailwind

Source: https://tailwindcss.com/docs/installation/framework-guides/meteor

Create a postcss.config.mjs file in the project root and register the @tailwindcss/postcss plugin.

```javascript
export default {
  plugins: {
    "@tailwindcss/postcss": {},
  },
};
```

--------------------------------

### Adding a Simple Custom Utility

Source: https://tailwindcss.com/docs/adding-custom-styles

Use the @utility directive to add a new utility class like 'content-auto' for CSS features not included by default.

```css
@utility content-auto {
  content-visibility: auto;
}
```

--------------------------------

### Responsive Block Size Design

Source: https://tailwindcss.com/docs/block-size

Prefix a `block-size` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="block-1/2 md:block-full ...">  <!-- ... --></div>
```

--------------------------------

### Apply clear utilities at specific breakpoints

Source: https://tailwindcss.com/docs/clear

Prefix a `clear` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<p class="clear-left md:clear-none ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Apply Styles Within a Breakpoint Range

Source: https://tailwindcss.com/docs/responsive-design

Use a responsive variant stacked with a `max-*` variant to apply a utility only within a specific breakpoint range. This is useful for fine-tuning layouts at certain screen sizes.

```html
<div class="md:max-xl:flex">  <!-- ... --></div>
```

--------------------------------

### Applying Predefined Background Position Utilities in HTML

Source: https://tailwindcss.com/docs/background-position

Use utilities like bg-center, bg-right, and bg-top-left to control the position of an element's background image.

```html
<div class="bg-[url(/img/mountains.jpg)] bg-top-left"></div><div class="bg-[url(/img/mountains.jpg)] bg-top"></div><div class="bg-[url(/img/mountains.jpg)] bg-top-right"></div><div class="bg-[url(/img/mountains.jpg)] bg-left"></div><div class="bg-[url(/img/mountains.jpg)] bg-center"></div><div class="bg-[url(/img/mountains.jpg)] bg-right"></div><div class="bg-[url(/img/mountains.jpg)] bg-bottom-left"></div><div class="bg-[url(/img/mountains.jpg)] bg-bottom"></div><div class="bg-[url(/img/mountains.jpg)] bg-bottom-right"></div>
```

--------------------------------

### Configure Tailwind CSS Vite plugin in Qwik

Source: https://tailwindcss.com/docs/installation/framework-guides/qwik

Add the `@tailwindcss/vite` plugin to your `vite.config.ts` file to enable Tailwind CSS processing.

```typescript
import { defineConfig } from 'vite'import { qwikVite } from "@builder.io/qwik/optimizer";import { qwikCity } from "@builder.io/qwik-city/vite";// …import tailwindcss from '@tailwindcss/vite'export default defineConfig(({ command, mode }): UserConfig => {
  return {
    plugins: [
      tailwindcss(),
      qwikCity(),
      qwikVite(),
      tsconfigPaths(),
    ],
    // …
  }
})
```

--------------------------------

### Enable Small Caps with font-feature-settings

Source: https://tailwindcss.com/docs/font-feature-settings

Use the `font-features-[<value>]` utility to enable OpenType features like small caps.

```html
<p class="font-features-['smcp'] ...">This text uses small caps.</p>
```

--------------------------------

### Applying Basic mask-clip Utilities in HTML

Source: https://tailwindcss.com/docs/mask-clip

Use `mask-clip-border`, `mask-clip-padding`, and `mask-clip-content` classes to define the bounding box for an element's mask.

```html
<div class="mask-clip-border border-3 p-1.5 mask-[url(/img/circle.png)] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-clip-padding border-3 p-1.5 mask-[url(/img/circle.png)] bg-[url(/img/mountains.jpg)] ..."></div><div class="mask-clip-content border-3 p-1.5 mask-[url(/img/circle.png)] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Responsive Line Height

Source: https://tailwindcss.com/docs/line-height

Utilize breakpoint prefixes like `md:` with line height utilities to apply them only at specific screen sizes and above. This enables responsive text spacing.

```html
<p class="leading-5 md:leading-6 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Use responsive breakpoints with scrollbar colors

Source: https://tailwindcss.com/docs/scrollbar-color

Prefix scrollbar color utilities with breakpoint variants like md: to apply different colors at specific screen sizes.

```html
<div class="scrollbar-thumb-slate-900 scrollbar-track-slate-200 md:scrollbar-thumb-sky-700 ...">  <!-- ... --></div>
```

--------------------------------

### Basic text-decoration-thickness utilities

Source: https://tailwindcss.com/docs/text-decoration-thickness

Use `decoration-<number>` utilities like `decoration-2` and `decoration-4` to change the text decoration thickness of an element.

```html
<p class="underline decoration-1">The quick brown fox...</p><p class="underline decoration-2">The quick brown fox...</p><p class="underline decoration-4">The quick brown fox...</p>
```

--------------------------------

### Custom min-width with CSS variables

Source: https://tailwindcss.com/docs/min-width

Use the min-w-(<custom-property>) syntax as shorthand for min-w-[var(<custom-property>)] to reference CSS custom properties.

```html
<div class="min-w-(--my-min-width) ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive Skew Design

Source: https://tailwindcss.com/docs/skew

Apply skew utilities responsively by prefixing them with a breakpoint variant, such as `md:skew-12`. This ensures the skew effect only applies at specific screen sizes and above.

```html
<img class="skew-3 md:skew-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive Field Sizing

Source: https://tailwindcss.com/docs/field-sizing

Apply `field-sizing` utilities conditionally at different breakpoints using responsive prefixes like `md:`. This allows for dynamic sizing adjustments based on screen size.

```html
<input class="field-sizing-content md:field-sizing-fixed ..." />
```

--------------------------------

### Arbitrary Max Height with calc()

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Utilize arbitrary values with CSS `calc()` for dynamic sizing, even with theme values.

```html
<div class="max-h-[calc(100dvh-(--spacing(6)))]">  <!-- ... --></div>
```

--------------------------------

### Responsive scroll-snap-align

Source: https://tailwindcss.com/docs/scroll-snap-align

Prefix a `scroll-snap-align` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="snap-center md:snap-start ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive justify-content

Source: https://tailwindcss.com/docs/justify-content

Prefix a `justify-content` utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<div class="flex justify-start md:justify-between ...">  <!-- ... --></div>
```

--------------------------------

### Import Preflight Styles

Source: https://tailwindcss.com/docs/preflight

Preflight is automatically injected into the `base` layer when `tailwindcss` is imported.

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/preflight.css" layer(base);
@import "tailwindcss/utilities.css" layer(utilities);
```

--------------------------------

### Responsive max-block-size Design

Source: https://tailwindcss.com/docs/max-block-size

Apply `max-block-size` utilities responsively by prefixing them with a breakpoint variant, like `md:max-block-screen`. This ensures the utility is only active at medium screen sizes and above.

```html
<div class="block-48 max-block-full md:max-block-screen ...">  <!-- ... --></div>
```

--------------------------------

### Basic Background Color Utilities

Source: https://tailwindcss.com/docs/background-color

Use predefined color utilities like `bg-white`, `bg-indigo-500`, and `bg-transparent` to set the background color of elements.

```html
<button class="bg-blue-500 ...">Button A</button><button class="bg-cyan-500 ...">Button B</button><button class="bg-pink-500 ...">Button C</button>
```

--------------------------------

### Using a thin scrollbar

Source: https://tailwindcss.com/docs/scrollbar-width

Use the `scrollbar-thin` utility to use a thinner scrollbar.

```html
<div class="scrollbar-thin overflow-auto ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Backdrop Sepia

Source: https://tailwindcss.com/docs/backdrop-filter-sepia

Prefix a backdrop-filter: sepia() utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="backdrop-sepia md:backdrop-sepia-0 ...">  <!-- ... --></div>
```

--------------------------------

### Apply Responsive Hue Rotation with Tailwind CSS

Source: https://tailwindcss.com/docs/filter-hue-rotate

Apply hue rotation utilities conditionally at different screen sizes using responsive variants like `md:`.

```html
<img class="hue-rotate-60 md:hue-rotate-0 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Background Size: Cover

Source: https://tailwindcss.com/docs/background-size

Use `bg-cover` to scale the background image to fill the element, cropping if necessary. Apply `bg-center` to center the image.

```html
<div class="bg-[url(/img/mountains.jpg)] bg-cover bg-center"></div>
```

--------------------------------

### Whitespace Pre Wrap Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-pre-wrap` utility to preserve newlines and spaces. Text will be wrapped normally.

```html
<p class="whitespace-pre-wrap">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Grid Container Layout

Source: https://tailwindcss.com/docs/display

Utilize the 'grid' utility to define a grid container for creating multi-column or multi-row layouts.

```html
<div class="grid grid-cols-3 grid-rows-3 gap-4">  <!-- ... --></div>
```

--------------------------------

### Responsive Backdrop Opacity

Source: https://tailwindcss.com/docs/backdrop-filter-opacity

Control backdrop filter opacity responsively by prefixing utilities with breakpoint variants like `md:`. This ensures the utility is applied only at the specified screen size and above.

```html
<div class="backdrop-opacity-100 md:backdrop-opacity-60 ...">  <!-- ... --></div>
```

--------------------------------

### Import Tailwind CSS in app.css

Source: https://tailwindcss.com/docs/installation/framework-guides/laravel/vite

Add an `@import` statement to your `app.css` file to include Tailwind CSS and specify directories for utility scanning.

```css
@import "tailwindcss";@source "../../vendor/laravel/framework/src/Illuminate/Pagination/resources/views/*.blade.php";@source "../../storage/framework/views/*.php";@source "../**/*.blade.php";@source "../**/*.js";
```

--------------------------------

### Responsive Caption Placement

Source: https://tailwindcss.com/docs/caption-side

Prefix a `caption-side` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above. This allows for adaptive table caption placement based on viewport size.

```html
<caption class="caption-top md:caption-bottom ...">  <!-- ... --></caption>
```

--------------------------------

### Setting a custom flex grow factor with Tailwind CSS `grow-[<value>]`

Source: https://tailwindcss.com/docs/flex-grow

Define a custom flex grow factor using arbitrary values within square brackets, such as `grow-[25vw]`.

```html
<div class="grow-[25vw] ...">  <!-- ... --></div>
```

--------------------------------

### Custom Transition Timing Function

Source: https://tailwindcss.com/docs/transition-timing-function

Set a custom transition timing function using the `ease-[<value>]` syntax for arbitrary cubic-bezier values. For CSS variables, use the `ease-(<custom-property>)` shorthand.

```html
<button class="ease-[cubic-bezier(0.95,0.05,0.795,0.035)] ...">  <!-- ... --></button>
```

```html
<button class="ease-(--my-ease) ...">  <!-- ... --></button>
```

--------------------------------

### Set Border-inline-start Color to Transparent

Source: https://tailwindcss.com/docs/border-color

Use `border-s-transparent` to make the inline-start border invisible.

```css
border-inline-start-color: transparent;
```

--------------------------------

### Configure Tailwind plugin in config.exs

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Set the Tailwind CSS version and configure input/output asset paths in your configuration file.

```elixir
config :tailwind,  version: "4.1.10",  myproject: [    args: ~w(      --input=assets/css/app.css      --output=priv/static/assets/app.css    ),    cd: Path.expand("..", __DIR__)  ]
```

--------------------------------

### Responsive justify-items

Source: https://tailwindcss.com/docs/justify-items

Prefix a `justify-items` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="grid justify-items-start md:justify-items-center ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Inline Size

Source: https://tailwindcss.com/docs/inline-size

Prefix an `inline-size` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="inline-1/2 md:inline-full ...">  <!-- ... --></div>
```

--------------------------------

### Custom Backdrop Sepia Value

Source: https://tailwindcss.com/docs/backdrop-filter-sepia

Use the backdrop-sepia-[<value>] syntax to set the backdrop sepia based on a completely custom value.

```html
<div class="backdrop-sepia-[.25] ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Text Shadow Design

Source: https://tailwindcss.com/docs/text-shadow

Prefix a `text-shadow` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<p class="text-shadow-none md:text-shadow-lg ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Responsive flex-direction

Source: https://tailwindcss.com/docs/flex-direction

Prefix a `flex-direction` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This allows for dynamic layout adjustments based on viewport size.

```html
<div class="flex flex-col md:flex-row ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Background Color

Source: https://tailwindcss.com/docs/background-color

Apply background colors responsively by prefixing a color utility with a breakpoint variant like `md:` to control the color at specific screen sizes and above.

```html
<div class="bg-blue-500 md:bg-green-500 ...">  <!-- ... --></div>
```

--------------------------------

### Set Gradient Color Stops

Source: https://tailwindcss.com/docs/background-image

Define the colors for gradient transitions using `from-<color>`, `via-<color>`, and `to-<color>` utilities. These utilities set the `--tw-gradient-from`, `--tw-gradient-via`, and `--tw-gradient-to` CSS variables.

```html
<div class="bg-linear-to-r from-indigo-500 via-purple-500 to-pink-500 ..."></div>
```

--------------------------------

### Set Custom Numeric Opacity with Tailwind CSS

Source: https://tailwindcss.com/docs/opacity

Illustrates using the arbitrary value syntax `opacity-[<value>]` to set a custom numeric opacity level.

```html
<button class="opacity-[.67] ...">  <!-- ... --></button>
```

--------------------------------

### Whitespace Pre Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-pre` utility to preserve newlines and spaces. Text will not be wrapped.

```html
<p class="overflow-auto whitespace-pre">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Applying Responsive transform-style in HTML

Source: https://tailwindcss.com/docs/transform-style

Use responsive prefixes like `md:` to apply `transform-style` utilities only at specific screen sizes and above, such as `md:transform-flat`.

```html
<div class="transform-3d md:transform-flat ...">  <!-- ... --></div>
```

--------------------------------

### Basic HTML Structure with Tailwind Classes

Source: https://tailwindcss.com/docs/installation/using-postcss

Reference your compiled CSS file in the HTML `<head>` and apply Tailwind utility classes directly to elements in your markup.

```HTML
<!doctype html><html><head>  <meta charset="UTF-8">  <meta name="viewport" content="width=device-width, initial-scale=1.0">  <link href="/dist/styles.css" rel="stylesheet"></head><body>  <h1 class="text-3xl font-bold underline">    Hello world!  </h1></body></html>
```

--------------------------------

### Adding Rings

Source: https://tailwindcss.com/docs/box-shadow

Apply a solid box-shadow ring to an element using `ring` or `ring-<number>` utilities like `ring-2` and `ring-4`. By default, rings match the element's `currentColor`.

```html
<button class="ring ...">Subscribe</button><button class="ring-2 ...">Subscribe</button><button class="ring-4 ...">Subscribe</button>
```

--------------------------------

### Custom backdrop brightness value with bracket notation

Source: https://tailwindcss.com/docs/backdrop-filter-brightness

Use the backdrop-brightness-[<value>] syntax to set backdrop brightness with a completely custom value not covered by predefined utilities.

```html
<div class="backdrop-brightness-[1.75] ...">  <!-- ... --></div>
```

--------------------------------

### Set Percentage-Based Maximum Width with Tailwind CSS

Source: https://tailwindcss.com/docs/max-width

Apply `max-w-full` or `max-w-<fraction>` utilities to set a percentage-based maximum width for an element.

```html
<div class="w-full max-w-9/10 ...">max-w-9/10</div><div class="w-full max-w-3/4 ...">max-w-3/4</div><div class="w-full max-w-1/2 ...">max-w-1/2</div><div class="w-full max-w-1/3 ...">max-w-1/3</div>
```

--------------------------------

### Custom Cursor with URL in HTML

Source: https://tailwindcss.com/docs/cursor

Use the "cursor-[<value>]" syntax to define a custom cursor using a URL to an image file, falling back to a generic pointer if the image fails.

```html
<button class="cursor-[url(hand.cur),_pointer] ...">  <!-- ... --></button>
```

--------------------------------

### Clear Start/End with Tailwind CSS (Logical Properties)

Source: https://tailwindcss.com/docs/clear

Utilize `clear-start` and `clear-end` for clearing behavior that adapts to text direction. `clear-start` maps to left in LTR and right in RTL, while `clear-end` maps to right in LTR and left in RTL.

```html
<article dir="rtl">
  <img class="float-left ..." src="/img/green-mountains.jpg" />
  <img class="float-right ..." src="/img/green-mountains.jpg" />
  <p class="clear-end ...">...ربما يمكننا العيش بدون مكتبات،</p>
</article>
```

--------------------------------

### Inline Sizes from Container Scale

Source: https://tailwindcss.com/docs/inline-size

Use utilities like `inline-sm` and `inline-xl` to set an element to a fixed inline size based on the container scale.

```html
<div class="inline-xl ...">inline-xl</div><div class="inline-lg ...">inline-lg</div><div class="inline-md ...">inline-md</div><div class="inline-sm ...">inline-sm</div><div class="inline-xs ...">inline-xs</div><div class="inline-2xs ...">inline-2xs</div><div class="inline-3xs ...">inline-3xs</div>
```

--------------------------------

### Undo line clamping with line-clamp-none

Source: https://tailwindcss.com/docs/line-clamp

Apply `line-clamp-none` to remove any previously applied line clamping, allowing text to display fully.

```html
<p class="line-clamp-3 lg:line-clamp-none">
  <!-- ... -->
</p>
```

--------------------------------

### Using stacked fractions

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `stacked-fractions` utility to replace numbers separated by a slash with common stacked fractions in fonts that support them.

```html
<p class="stacked-fractions ...">1/2 3/4 5/6</p>
```

--------------------------------

### Responsive list-style-type with breakpoint variant

Source: https://tailwindcss.com/docs/list-style-type

Prefix a list-style-type utility with a breakpoint variant like md: to apply the style only at medium screen sizes and above.

```html
<ul class="list-none md:list-disc ...">
  <!-- ... -->
</ul>
```

--------------------------------

### Add Tailwind CSS to postcss.config.mjs

Source: https://tailwindcss.com/docs/installation/framework-guides/emberjs

Create `postcss.config.mjs` and add the `@tailwindcss/postcss` plugin to the PostCSS configuration.

```javascript
export default {  plugins: {    "@tailwindcss/postcss": {},  }}
```

--------------------------------

### Using arbitrary values for text-decoration-color

Source: https://tailwindcss.com/docs/text-decoration-color

Set arbitrary values for text decoration colors directly using the `decoration-[<value>]` utility. This is useful for applying specific colors not covered by predefined classes.

```html
<p class="decoration-[rgb(255_0_0)]">Red Text Decoration</p>
<p class="decoration-[oklch(60%_0.2_130)]">Oklch Text Decoration</p>
```

--------------------------------

### Configure PostCSS plugins for Tailwind CSS

Source: https://tailwindcss.com/docs/installation/framework-guides/rspack/vue

Create "postcss.config.mjs" to add the "@tailwindcss/postcss" plugin.

```JavaScript
export default {  plugins: {    "@tailwindcss/postcss": {}  }};
```

--------------------------------

### Translate elements using percentages

Source: https://tailwindcss.com/docs/translate

Apply `translate-<fraction>` utilities to translate an element on both X and Y axes by a percentage of its own size.

```html
<img class="-translate-1/4 ..." src="/img/mountains.jpg" /><img class="translate-1/6 ..." src="/img/mountains.jpg" /><img class="translate-1/2 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Set Border-inline-start Color with White

Source: https://tailwindcss.com/docs/border-color

Use `border-s-white` to set the inline-start border color to white using the predefined `--color-white` variable.

```css
border-inline-start-color: var(--color-white); /* #fff */
```

--------------------------------

### Basic Backdrop Opacity Utilities

Source: https://tailwindcss.com/docs/backdrop-filter-opacity

Use utilities like `backdrop-opacity-10`, `backdrop-opacity-60`, and `backdrop-opacity-95` to control the opacity of backdrop filters. These classes apply specific percentage values to the backdrop-filter opacity property.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert backdrop-opacity-10 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert backdrop-opacity-60 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-invert backdrop-opacity-95 ..."></div></div>
```

--------------------------------

### Responsive Outline Width

Source: https://tailwindcss.com/docs/outline-width

Apply outline width utilities at specific breakpoints using prefix variants like `md:`. This ensures the outline width changes based on screen size.

```html
<div class="outline md:outline-2 ...">  <!-- ... --></div>
```

--------------------------------

### Responsive design for forced colors

Source: https://tailwindcss.com/docs/forced-color-adjust

Prefix a `forced-color-adjust` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="forced-color-adjust-none md:forced-color-adjust-auto ...">
  <!-- ... -->
</div>
```

--------------------------------

### Distribute content with space around in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-around` to distribute grid items with equal space around each row and column along both axes.

```html
<div class="grid h-48 grid-cols-2 place-content-around gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Responsive list-style-position

Source: https://tailwindcss.com/docs/list-style-position

Apply `list-style-position` utilities conditionally at different screen sizes using breakpoint prefixes like `md:`.

```html
<ul class="list-outside md:list-inside ...">  <!-- ... --></ul>
```

--------------------------------

### Basic list-style-position

Source: https://tailwindcss.com/docs/list-style-position

Use `list-inside` to position list markers within the text flow and `list-outside` to position them outside.

```html
<ul class="list-inside">  <li>5 cups chopped Porcini mushrooms</li>  <!-- ... --></ul><ul class="list-outside">  <li>5 cups chopped Porcini mushrooms</li>  <!-- ... --></ul>
```

--------------------------------

### Custom Tab Size with `tab-[<value>]`

Source: https://tailwindcss.com/docs/tab-size

Use the `tab-[<value>]` syntax to set the tab size based on a completely custom value, such as pixels.

```html
<pre class="tab-[12px] ...">  <!-- ... --></pre>
```

--------------------------------

### Responsive Sepia Filter

Source: https://tailwindcss.com/docs/filter-sepia

Apply sepia filters conditionally at different screen sizes using responsive prefixes like `md:`.

```html
<img class="sepia md:sepia-0 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Apply Responsive Text Color with Tailwind CSS Breakpoints

Source: https://tailwindcss.com/docs/color

Prefix a color utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<p class="text-blue-600 md:text-green-600 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Basic max-inline-size Utilities

Source: https://tailwindcss.com/docs/max-inline-size

Use `max-inline-<number>` utilities to set a fixed maximum inline size based on the spacing scale. Apply these utilities to elements to control their width within the defined scale.

```html
<div class="inline-full max-inline-96 ...">max-inline-96</div><div class="inline-full max-inline-80 ...">max-inline-80</div><div class="inline-full max-inline-64 ...">max-inline-64</div><div class="inline-full max-inline-48 ...">max-inline-48</div><div class="inline-full max-inline-40 ...">max-inline-40</div><div class="inline-full max-inline-32 ...">max-inline-32</div>
```

--------------------------------

### Match Viewport Block Size

Source: https://tailwindcss.com/docs/block-size

Use the `block-screen` utility to make an element span the entire block size of the viewport.

```html
<div class="block-screen">  <!-- ... --></div>
```

--------------------------------

### Apply Responsive Gaps with Tailwind CSS Breakpoints

Source: https://tailwindcss.com/docs/gap

Prefix `gap` utilities with breakpoint variants like `md:` to apply different gap sizes based on screen dimensions, such as `md:gap-6` for medium screens and above.

```html
<div class="grid gap-4 md:gap-6 ...">  <!-- ... --></div>
```

--------------------------------

### Enable Resizing in All Directions

Source: https://tailwindcss.com/docs/resize

Use the `resize` utility to allow an element to be resized both horizontally and vertically. This is useful for elements like textareas where user-adjustable sizing is desired.

```html
<textarea class="resize rounded-md ..."></textarea>
```

--------------------------------

### Basic Box Shadows

Source: https://tailwindcss.com/docs/box-shadow

Apply different sized outer box shadows using utilities like `shadow-sm`, `shadow-md`, `shadow-lg`, and `shadow-xl`. These utilities control the overall size and spread of the shadow.

```html
<div class="shadow-md ..."></div><div class="shadow-lg ..."></div><div class="shadow-xl ..."></div>
```

--------------------------------

### Responsive Margin Design

Source: https://tailwindcss.com/docs/margin

Prefix a margin utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<div class="mt-4 md:mt-8 ...">  <!-- ... --></div>
```

--------------------------------

### Predefined Background Colors (Mist Scale)

Source: https://tailwindcss.com/docs/background-color

Applies background colors from the mist color scale. Each class maps to a specific CSS variable.

```html
<div class="bg-mist-50">...</div>
<div class="bg-mist-100">...</div>
<div class="bg-mist-200">...</div>
<div class="bg-mist-300">...</div>
<div class="bg-mist-400">...</div>
<div class="bg-mist-500">...</div>
<div class="bg-mist-600">...</div>
<div class="bg-mist-700">...</div>
<div class="bg-mist-800">...</div>
<div class="bg-mist-900">...</div>
<div class="bg-mist-950">...</div>
```

--------------------------------

### Responsive font-family utility

Source: https://tailwindcss.com/docs/font-family

Apply font-family utilities responsively by prefixing them with a breakpoint variant like 'md:'.

```html
<p class="font-sans md:font-serif ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Matching Theme Values with Functional Utilities

Source: https://tailwindcss.com/docs/adding-custom-styles

Use the `--value(--theme-key-*)` syntax to resolve utility values against a predefined set of theme keys.

```css
@theme {
  --tab-size-2: 2;
  --tab-size-4: 4;
  --tab-size-github: 8;
}
@utility tab-* {
  tab-size: --value(--tab-size-*);
}
```

--------------------------------

### Responsive justify-self

Source: https://tailwindcss.com/docs/justify-self

Prefix a `justify-self` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="justify-self-start md:justify-self-end ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Width with Tailwind CSS Breakpoint Variants

Source: https://tailwindcss.com/docs/width

Apply width utilities conditionally at different screen sizes by prefixing them with responsive variants like `md:`.

```html
<div class="w-1/2 md:w-full ...">  <!-- ... --></div>
```

--------------------------------

### Set Border-inline-start Color to Inherit

Source: https://tailwindcss.com/docs/border-color

Use `border-s-inherit` to make the inline-start border color inherit from its parent.

```css
border-inline-start-color: inherit;
```

--------------------------------

### Configure Container Utility in v4

Source: https://tailwindcss.com/docs/upgrade-guide

Extend the `container` utility using the `@utility` directive to customize its properties like margin and padding in v4.

```css
@utility container {
  margin-inline: auto;
  padding-inline: 2rem;
}
```

--------------------------------

### Apply Conditional Opacity with Tailwind CSS

Source: https://tailwindcss.com/docs/opacity

Shows how to use variant prefixes like `disabled:` to apply opacity utilities conditionally based on element state.

```html
<input class="opacity-100 disabled:opacity-75 ..." type="text" />
```

--------------------------------

### Applying Custom Container Sizes

Source: https://tailwindcss.com/docs/responsive-design

Apply custom container sizes in your HTML markup using the defined variants, such as `@8xl` for a custom 8xl breakpoint.

```html
<div class="@container">  <div class="flex flex-col @8xl:flex-row">    <!-- ... -->  </div></div>
```

--------------------------------

### Applying Utility Classes with Multi-Cursor Editing

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Use multi-cursor editing to efficiently apply or modify utility classes across multiple elements in a single file when duplication is localized.

```html
<nav class="flex justify-center space-x-4">  <a href="/dashboard" class="font-medium rounded-lg px-3 py-2 text-gray-700 hover:bg-gray-100 hover:text-gray-900">    Home  </a>  <a href="/team" class="font-medium rounded-lg px-3 py-2 text-gray-700 hover:bg-gray-100 hover:text-gray-900">    Team  </a>  <a href="/projects" class="font-medium rounded-lg px-3 py-2 text-gray-700 hover:bg-gray-100 hover:text-gray-900">    Projects  </a>  <a href="/reports" class="font-medium rounded-lg px-3 py-2 text-gray-700 hover:bg-gray-100 hover:text-gray-900">    Reports  </a></nav>
```

--------------------------------

### Block and Inline Element Display

Source: https://tailwindcss.com/docs/display

Use 'inline', 'inline-block', and 'block' utilities to control how elements flow with text. 'inline' wraps normally, 'inline-block' wraps the element to its parent, and 'block' places the element on its own line.

```html
<p>  When controlling the flow of text, using the CSS property <span class="inline">display: inline</span> will cause the  text inside the element to wrap normally.</p><p>  While using the property <span class="inline-block">display: inline-block</span> will wrap the element to prevent the  text inside from extending beyond its parent.</p><p>  Lastly, using the property <span class="block">display: block</span> will put the element on its own line and fill its  parent.</p>
```

--------------------------------

### Logical Margin Properties (Block)

Source: https://tailwindcss.com/docs/margin

Use the `mbs-<number>` and `mbe-<number>` utilities to set the `margin-block-start` and `margin-block-end` logical properties, which map to top or bottom based on writing mode.

```html
<div class="mbs-8 ...">mbs-8</div>
```

--------------------------------

### Adding Default Base Styles

Source: https://tailwindcss.com/docs/adding-custom-styles

Set default styles for the entire page by applying classes to the html or body elements.

```html
<!doctype html>
<html lang="en" class="bg-gray-100 font-serif text-gray-900">
  <!-- ... -->
</html>
```

--------------------------------

### Responsive Text Alignment

Source: https://tailwindcss.com/docs/text-align

Apply text alignment utilities conditionally at different screen sizes using breakpoint prefixes like `md:`. This allows for responsive text alignment strategies.

```html
<p class="text-left md:text-center ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Logical Margin Properties (Inline)

Source: https://tailwindcss.com/docs/margin

Use `ms-<number>` or `me-<number>` utilities to set the `margin-inline-start` and `margin-inline-end` logical properties, which adapt to the text direction.

```html
<div>  <div dir="ltr">    <div class="ms-8 ...">ms-8</div>    <div class="me-8 ...">me-8</div>  </div>  <div dir="rtl">    <div class="ms-8 ...">ms-8</div>    <div class="me-8 ...">me-8</div>  </div></div>
```

--------------------------------

### Using Size Containers

Source: https://tailwindcss.com/docs/responsive-design

Mark an element as a size container to use block-size dependent container query units like `cqb`. This is useful when the layout depends on the block dimension.

```html
<div class="@container-size">  <div class="h-[50cqb]">    <!-- ... -->  </div></div>
```

--------------------------------

### Responsive Mask Image

Source: https://tailwindcss.com/docs/mask-image

Apply mask image utilities responsively by prefixing them with a breakpoint variant like `md:`. This allows the mask to change at different screen sizes.

```html
<div class="mask-radial-from-70% md:mask-radial-from-50% ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Scroll Snapping

Source: https://tailwindcss.com/docs/scroll-snap-type

Apply `scroll-snap-type` utilities at specific breakpoints using responsive prefixes like `md:`.

```html
<div class="snap-none md:snap-x ...">
  <!-- ... -->
</div>
```

--------------------------------

### Define a custom breakpoint theme variable

Source: https://tailwindcss.com/docs/theme

Define a new responsive breakpoint using the `--breakpoint-*` namespace to create a corresponding variant for conditional styling.

```css
@import "tailwindcss";@theme {  --breakpoint-3xl: 120rem;}
```

--------------------------------

### Targeting Dark Mode

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Use the 'dark:' prefix to apply styles specifically when dark mode is enabled. This requires separate classes for light and dark mode.

```html
<div class="bg-white dark:bg-gray-800 rounded-lg px-6 py-8 ring shadow-xl ring-gray-900/5">
  <div>
    <span class="inline-flex items-center justify-center rounded-md bg-indigo-500 p-2 shadow-lg">
      <svg
        class="h-6 w-6 text-white"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        aria-hidden="true"
      >
        <!-- ... -->
      </svg>
    </span>
  </div>
  <h3 class="text-gray-900 dark:text-white mt-5 text-base font-medium tracking-tight ">Writes upside-down</h3>
  <p class="text-gray-500 dark:text-gray-400 mt-2 text-sm ">
    The Zero Gravity Pen can be used to write in any orientation, including upside-down. It even works in outer space.
  </p>
</div>
```

--------------------------------

### Custom value min-inline-size utility

Source: https://tailwindcss.com/docs/min-inline-size

Use the `min-inline-[<value>]` syntax for completely custom minimum inline sizes, such as `min-inline-[220px]`.

```html
<div class="min-inline-[220px] ...">  <!-- ... --></div>
```

--------------------------------

### Responsive z-index

Source: https://tailwindcss.com/docs/z-index

Prefix a z-index utility with a breakpoint variant like md: to apply it at specific screen sizes and above.

```html
<div class="z-0 md:z-50 ...">  <!-- ... --></div>
```

--------------------------------

### Stacking Variants for Multiple Conditions

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Demonstrates stacking variants like 'disabled:' and 'hover:' to apply styles only when multiple conditions are met simultaneously.

```html
<button class="bg-sky-500 disabled:hover:bg-sky-500 ...">Save changes</button>
```

--------------------------------

### Inset Rings

Source: https://tailwindcss.com/docs/box-shadow

Apply a solid inset box-shadow ring using `inset-ring` or `inset-ring-<number>` utilities like `inset-ring-2` and `inset-ring-4`. These match the element's `currentColor` by default.

```html
<button class="inset-ring ...">Subscribe</button><button class="inset-ring-2 ...">Subscribe</button><button class="inset-ring-4 ...">Subscribe</button>
```

--------------------------------

### Transition with custom height property

Source: https://tailwindcss.com/docs/transition-property

Use the `transition-[<value>]` syntax to set transition properties based on a custom value, such as 'height'.

```html
<button class="transition-[height] ...">  <!-- ... --></button>
```

--------------------------------

### Use Arbitrary Breakpoint Values

Source: https://tailwindcss.com/docs/responsive-design

For one-off breakpoints not suitable for the theme, use `min` or `max` variants with arbitrary values directly in your markup.

```html
<div class="max-[600px]:bg-sky-300 min-[320px]:text-center">  <!-- ... --></div>
```

--------------------------------

### Responsive font-feature-settings

Source: https://tailwindcss.com/docs/font-feature-settings

Prefix a `font-feature-settings` utility with a breakpoint variant like `md:` to apply it only at medium screen sizes and above.

```html
<p class="font-features-['tnum'] md:font-features-['smcp'] ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Set minimum block size with CSS variable

Source: https://tailwindcss.com/docs/min-block-size

Use the `min-block-(<custom-property>)` syntax as a shorthand for `min-block-[var(<custom-property>)]` to set the minimum block size using a CSS variable.

```html
<div class="min-block-(--my-min-block-size) ...">  <!-- ... --></div>
```

--------------------------------

### Set Both Width and Height with Tailwind CSS

Source: https://tailwindcss.com/docs/height

Use `size-<value>` utilities like `size-16` to simultaneously set both the width and height of an element.

```html
<div class="size-16 ...">size-16</div><div class="size-20 ...">size-20</div><div class="size-24 ...">size-24</div><div class="size-32 ...">size-32</div><div class="size-40 ...">size-40</div>
```

--------------------------------

### Basic Backdrop Contrast Filter

Source: https://tailwindcss.com/docs/backdrop-filter-contrast

Use utilities like `backdrop-contrast-50` and `backdrop-contrast-200` to control an element's backdrop contrast. These apply `backdrop-filter: contrast(<number>%).`

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-contrast-50 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-contrast-200 ..."></div></div>
```

--------------------------------

### Apply custom padding with CSS variables

Source: https://tailwindcss.com/docs/padding

Use the `p-(<custom-property>)` syntax to apply padding using a CSS variable.

```html
<div class="p-(--my-padding) ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Zoom

Source: https://tailwindcss.com/docs/zoom

Prefix a `zoom` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This enables responsive scaling.

```html
<div class="zoom-100 md:zoom-125 ...">  <!-- ... --></div>
```

--------------------------------

### Import global.css in gatsby-browser.js

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Create or update `gatsby-browser.js` at the root of your project to import the global CSS file.

```JavaScript
import './src/styles/global.css';
```

--------------------------------

### Updating Arbitrary Values in Grid and Object-Position Utilities

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, underscores `_` must be used instead of commas `,` to represent spaces in arbitrary values for `grid-cols-*`, `grid-rows-*`, and `object-*` utilities.

```html
<div class="grid-cols-[max-content,auto]"></div>
<div class="grid-cols-[max-content_auto]"></div>
```

--------------------------------

### Safelisting utilities with variants

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Generate classes with variants like hover and focus by including them in the `@source inline()` input.

```css
@import "tailwindcss";@source inline("{hover:,focus:,}underline");
```

--------------------------------

### Custom transition duration with bracket notation

Source: https://tailwindcss.com/docs/transition-duration

Use the duration-[<value>] syntax to apply arbitrary custom duration values not covered by standard utility classes.

```html
<button class="duration-[1s,15s] ...">  <!-- ... --></button>
```

--------------------------------

### Responsive backdrop filter with breakpoint variant

Source: https://tailwindcss.com/docs/backdrop-filter

Prefix a backdrop-filter utility with a breakpoint variant like md: to apply the utility only at medium screen sizes and above.

```html
<div class="backdrop-blur-sm md:backdrop-filter-none ..."></div>
```

--------------------------------

### Collapse table rows/columns with `collapse`

Source: https://tailwindcss.com/docs/visibility

Use the `collapse` utility to hide table rows, row groups, columns, and column groups as if they were set to `display: none`, but without impacting the size of other rows and columns. This makes it possible to dynamically toggle rows and columns without affecting the table layout.

```html
<table>  <thead>    <tr>      <th>Invoice #</th>      <th>Client</th>      <th>Amount</th>    </tr>  </thead>  <tbody>    <tr>      <td>#100</td>      <td>Pendant Publishing</td>      <td>$2,000.00</td>    </tr>    <tr class="collapse">      <td>#101</td>      <td>Kruger Industrial Smoothing</td>      <td>$545.00</td>    </tr>    <tr>      <td>#102</td>      <td>J. Peterman</td>      <td>$10,000.25</td>    </tr>  </tbody></table>
```

--------------------------------

### Flex item grow and shrink, respecting initial size

Source: https://tailwindcss.com/docs/flex

Use `flex-auto` to allow a flex item to grow and shrink, taking into account its initial size.

```html
<div class="flex ...">  <div class="w-14 flex-none ...">01</div>  <div class="w-64 flex-auto ...">02</div>  <div class="w-32 flex-auto ...">03</div></div>
```

--------------------------------

### Make element visible with `visible`

Source: https://tailwindcss.com/docs/visibility

Use the `visible` utility to make an element visible. This is mostly useful for undoing the `invisible` utility at different screen sizes.

```html
<div class="grid grid-cols-3 gap-4">  <div>01</div>  <div class="visible ...">02</div>  <div>03</div></div>
```

--------------------------------

### Transitioning Individual Transform Properties

Source: https://tailwindcss.com/docs/upgrade-guide

When customizing transitioned properties, include individual transform properties like `scale` instead of the general `transform` to ensure transitions work correctly.

```html
<button class="transition-[opacity,transform] hover:scale-150"></button><button class="transition-[opacity,scale] hover:scale-150"></button>
```

--------------------------------

### Customize blur utility values in theme

Source: https://tailwindcss.com/docs/filter-blur

Define custom blur sizes by setting --blur-* CSS variables within the @theme block to create new utility classes.

```css
@theme {  --blur-2xs: 2px; }
```

--------------------------------

### Basic Margin Utilities

Source: https://tailwindcss.com/docs/margin

Use `m-<number>` utilities like `m-4` and `m-8` to control the margin on all sides of an element.

```html
<div class="m-8 ...">m-8</div>
```

--------------------------------

### Setting Explicit Base Path for Source Detection

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Use the source() function when importing Tailwind to set an explicit base path for source detection. This is helpful in monorepos.

```css
@import "tailwindcss" source("../src");
```

--------------------------------

### Using Arbitrary Opacity Values in Tailwind CSS

Source: https://tailwindcss.com/docs/colors

Illustrates the use of arbitrary values and CSS variables for setting color opacity in Tailwind CSS classes.

```html
<div class="bg-pink-500/[71.37%]"><!-- ... --></div><div class="bg-cyan-400/(--my-alpha-value)"><!-- ... --></div>
```

--------------------------------

### Inline Grid Container Layout

Source: https://tailwindcss.com/docs/display

Employ the 'inline-grid' utility to create a grid container that flows inline with text.

```html
<span class="inline-grid grid-cols-3 gap-4">  <span>01</span>  <span>02</span>  <span>03</span>  <span>04</span>  <span>05</span>  <span>06</span></span><span class="inline-grid grid-cols-3 gap-4">  <span>01</span>  <span>02</span>  <span>03</span>  <span>04</span>  <span>05</span>  <span>06</span></span>
```

--------------------------------

### Use custom theme border radius

Source: https://tailwindcss.com/docs/border-radius

Apply custom border radius utilities defined in the theme configuration.

```html
<div class="rounded-5xl">  <!-- ... --></div>
```

--------------------------------

### Apply Responsive Outline Styles with Tailwind CSS

Source: https://tailwindcss.com/docs/outline-style

Prefix `outline-style` utilities with breakpoint variants like `md:` to apply styles conditionally based on screen size.

```html
<div class="outline md:outline-dashed ...">  <!-- ... --></div>
```

--------------------------------

### Style Details Content with :details-content Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles to the content of a <details> element when it is open.

```html
<details class="details-content:bg-gray-100 ...">
  <summary>Details</summary>
  This is a secret.
</details>
```

--------------------------------

### Set Uniform Gaps in Grid and Flexbox with Tailwind CSS

Source: https://tailwindcss.com/docs/gap

Use `gap-<number>` utilities like `gap-2` and `gap-4` to apply uniform spacing between both rows and columns in grid and flexbox layouts.

```html
<div class="grid grid-cols-2 gap-4">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Inline Flex Container Layout

Source: https://tailwindcss.com/docs/display

Create an inline flex container that integrates with surrounding text using the 'inline-flex' utility.

```html
<p>  Today I spent most of the day researching ways to ...  <span class="inline-flex items-baseline">    <img src="/img/kramer.jpg" class="mx-1 size-5 self-center rounded-full" />    <span>Kramer</span>  </span>  keeps telling me there is no way to make it work, that ...</p>
```

--------------------------------

### Basic list-style-image

Source: https://tailwindcss.com/docs/list-style-image

Use the `list-image-[<value>]` syntax to set a custom marker image for list items.

```html
<ul class="list-image-[url(/img/checkmark.png)]">  <li>5 cups chopped Porcini mushrooms</li>  <!-- ... --></ul>
```

--------------------------------

### Transitioning Outline Color

Source: https://tailwindcss.com/docs/upgrade-guide

The `transition` and `transition-colors` utilities now include `outline-color`. To prevent unwanted transitions, set the outline color unconditionally or for both states.

```html
<button class="transition hover:outline-2 hover:outline-cyan-500"></button>
<button class="outline-cyan-500 transition hover:outline-2"></button>
```

--------------------------------

### Horizontal Scroll Snapping

Source: https://tailwindcss.com/docs/scroll-snap-type

Use the `snap-x` utility to enable horizontal scroll snapping. Ensure children have snap alignment.

```html
<div class="snap-x ...">
  <div class="snap-center ...">
    <img src="/img/vacation-01.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-02.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-03.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-04.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-05.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-06.jpg" />
  </div>
</div>
```

--------------------------------

### Full viewport width

Source: https://tailwindcss.com/docs/width

Use w-screen to make an element span the entire width of the viewport.

```html
<div class="w-screen">  <!-- ... --></div>
```

--------------------------------

### Basic Backdrop Sepia

Source: https://tailwindcss.com/docs/backdrop-filter-sepia

Use backdrop-sepia, backdrop-sepia-50, and backdrop-sepia-0 to control the sepia effect applied to an element's backdrop.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-sepia-0 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-sepia-50 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-sepia ..."></div></div>
```

--------------------------------

### Apply Different Sized Text Shadows

Source: https://tailwindcss.com/docs/text-shadow

Use utilities like `text-shadow-sm` and `shadow-lg` to apply different sized text shadows to a text element.

```html
<p class="text-shadow-2xs ...">The quick brown fox...</p><p class="text-shadow-xs ...">The quick brown fox...</p><p class="text-shadow-sm ...">The quick brown fox...</p><p class="text-shadow-md ...">The quick brown fox...</p><p class="text-shadow-lg ...">The quick brown fox...</p>
```

--------------------------------

### Custom line clamping with line-clamp-[value]

Source: https://tailwindcss.com/docs/line-clamp

Use the `line-clamp-[<value>]` syntax to set the number of lines based on a custom value, such as a CSS variable calculation.

```html
<p class="line-clamp-[calc(var(--characters)/100)] ...">
  Lorem ipsum dolor sit amet...
</p>
```

--------------------------------

### Basic Sepia Filters

Source: https://tailwindcss.com/docs/filter-sepia

Apply predefined sepia filter levels to an element using `sepia-0`, `sepia-50`, or the default `sepia` class.

```html
<img class="sepia-0" src="/img/mountains.jpg" /><img class="sepia-50" src="/img/mountains.jpg" /><img class="sepia" src="/img/mountains.jpg" />
```

--------------------------------

### Using Contents Utility

Source: https://tailwindcss.com/docs/display

Use the `contents` utility to create a 'phantom' container whose children act like direct children of the parent. This is useful for complex flexbox layouts where intermediate elements should not affect the layout.

```html
<div class="flex ...">  <div class="flex-1 ...">01</div>  <div class="contents">    <div class="flex-1 ...">02</div>    <div class="flex-1 ...">03</div>  </div>  <div class="flex-1 ...">04</div></div>
```

--------------------------------

### Applying box-decoration-break responsively with Tailwind CSS

Source: https://tailwindcss.com/docs/box-decoration-break

Apply `box-decoration-break` utilities conditionally based on screen size by prefixing them with a breakpoint variant like `md:`.

```html
<div class="box-decoration-clone md:box-decoration-slice ...">  <!-- ... --></div>
```

--------------------------------

### Set minimum block size with custom value

Source: https://tailwindcss.com/docs/min-block-size

Use the `min-block-[<value>]` syntax to set the minimum block size based on a completely custom pixel value.

```html
<div class="min-block-[220px] ...">  <!-- ... --></div>
```

--------------------------------

### Basic mask-repeat

Source: https://tailwindcss.com/docs/mask-repeat

Use the `mask-repeat` utility to repeat the mask image both vertically and horizontally. This is useful for tiling mask images.

```html
<div class="mask-repeat mask-[url(/img/circle.png)] mask-size-[50px_50px] bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Basic Invert Filters

Source: https://tailwindcss.com/docs/filter-invert

Apply basic color inversion to an element using `invert-0`, `invert-20`, or the default `invert` utility.

```html
<img class="invert-0" src="/img/mountains.jpg" /><img class="invert-20" src="/img/mountains.jpg" /><img class="invert" src="/img/mountains.jpg" />
```

--------------------------------

### Import global styles as reference in CSS modules

Source: https://tailwindcss.com/docs/compatibility

Use @reference to import global styles in CSS modules so that @apply and theme variables are available. This ensures Tailwind processes the module with access to your theme context.

```CSS
@reference "../app.css";
button {
  @apply bg-blue-500;
}
```

--------------------------------

### Add Conic Gradient

Source: https://tailwindcss.com/docs/background-image

Utilize `bg-conic` and `bg-conic-<angle>` utilities, along with color stop utilities, to implement conic gradients. The `from-` and `via-` utilities specify the gradient colors.

```html
<div class="size-24 rounded-full bg-conic from-blue-600 to-sky-400 to-50%"></div><div class="size-24 rounded-full bg-conic-180 from-indigo-600 via-indigo-50 to-indigo-600"></div><div class="size-24 rounded-full bg-conic/decreasing from-violet-700 via-lime-300 to-violet-700"></div>
```

--------------------------------

### Using custom font-family utility

Source: https://tailwindcss.com/docs/font-family

Apply a custom font family utility defined in the theme.

```html
<div class="font-display">  <!-- ... --></div>
```

--------------------------------

### Style all descendants with the ** variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `**` variant to apply styles to all descendants, not just direct children. This is particularly useful when combined with other variants for more specific selection.

```html
<ul class="**:data-avatar:size-12 **:data-avatar:rounded-full ...">
  {#each items as item}
    <li>
      <img src={item.src} data-avatar />
      <p>{item.name}</p>
    </li>
  {/each}
</ul>
```

--------------------------------

### justify-normal

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-normal` utility to pack content items in their default position as if no `justify-content` value was set.

```html
<div class="flex justify-normal ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Block-start border color utilities

Source: https://tailwindcss.com/docs/border-color

Use these utilities to set the `border-block-start-color` property for an element. These classes map directly to Tailwind's color palette.

```html
<div class="border-bs-pink-600">...</div>
<div class="border-bs-pink-700">...</div>
<div class="border-bs-pink-800">...</div>
<div class="border-bs-pink-900">...</div>
<div class="border-bs-pink-950">...</div>
<div class="border-bs-rose-50">...</div>
<div class="border-bs-rose-100">...</div>
<div class="border-bs-rose-200">...</div>
<div class="border-bs-rose-300">...</div>
<div class="border-bs-rose-400">...</div>
<div class="border-bs-rose-500">...</div>
<div class="border-bs-rose-600">...</div>
<div class="border-bs-rose-700">...</div>
<div class="border-bs-rose-800">...</div>
<div class="border-bs-rose-900">...</div>
<div class="border-bs-rose-950">...</div>
<div class="border-bs-slate-50">...</div>
<div class="border-bs-slate-100">...</div>
<div class="border-bs-slate-200">...</div>
<div class="border-bs-slate-300">...</div>
<div class="border-bs-slate-400">...</div>
<div class="border-bs-slate-500">...</div>
<div class="border-bs-slate-600">...</div>
<div class="border-bs-slate-700">...</div>
<div class="border-bs-slate-800">...</div>
<div class="border-bs-slate-900">...</div>
<div class="border-bs-slate-950">...</div>
<div class="border-bs-gray-50">...</div>
<div class="border-bs-gray-100">...</div>
<div class="border-bs-gray-200">...</div>
<div class="border-bs-gray-300">...</div>
<div class="border-bs-gray-400">...</div>
<div class="border-bs-gray-500">...</div>
<div class="border-bs-gray-600">...</div>
<div class="border-bs-gray-700">...</div>
<div class="border-bs-gray-800">...</div>
<div class="border-bs-gray-900">...</div>
<div class="border-bs-gray-950">...</div>
<div class="border-bs-zinc-50">...</div>
<div class="border-bs-zinc-100">...</div>
<div class="border-bs-zinc-200">...</div>
<div class="border-bs-zinc-300">...</div>
<div class="border-bs-zinc-400">...</div>
<div class="border-bs-zinc-500">...</div>
<div class="border-bs-zinc-600">...</div>
<div class="border-bs-zinc-700">...</div>
```

--------------------------------

### Implementing a subgrid with grid-rows-subgrid

Source: https://tailwindcss.com/docs/grid-template-rows

Use grid-rows-subgrid to adopt the row tracks defined by the parent grid. Nested grid items can then use row-span and row-start utilities to position themselves within inherited row tracks.

```html
<div class="grid grid-flow-col grid-rows-4 gap-4">
  <div>01</div>
  <!-- ... -->
  <div>05</div>
  <div class="row-span-3 grid grid-rows-subgrid gap-4">
    <div class="row-start-2">06</div>
  </div>
  <div>07</div>
  <!-- ... -->
  <div>10</div>
</div>
```

--------------------------------

### Custom Contrast Value

Source: https://tailwindcss.com/docs/filter-contrast

Set a custom contrast value using the `contrast-[<value>]` syntax for precise control.

```html
<img class="contrast-[.25] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Basic object-position utilities

Source: https://tailwindcss.com/docs/object-position

Use utilities like object-left and object-bottom-right to specify how a replaced element's content should be positioned within its container. These utilities map directly to the object-position CSS property.

```html
<img class="size-24 object-top-left ..." src="/img/mountains.jpg" /><img class="size-24 object-top ..." src="/img/mountains.jpg" /><img class="size-24 object-top-right ..." src="/img/mountains.jpg" /><img class="size-24 object-left ..." src="/img/mountains.jpg" /><img class="size-24 object-center ..." src="/img/mountains.jpg" /><img class="size-24 object-right ..." src="/img/mountains.jpg" /><img class="size-24 object-bottom-left ..." src="/img/mountains.jpg" /><img class="size-24 object-bottom ..." src="/img/mountains.jpg" /><img class="size-24 object-bottom-right ..." src="/img/mountains.jpg" />
```

--------------------------------

### Extend Preflight with Custom Base Styles

Source: https://tailwindcss.com/docs/preflight

Add custom base styles on top of Preflight by including them in the `base` CSS layer. This allows for easy customization of default element styles.

```css
@layer base {
  h1 {
    font-size: var(--text-2xl);
  }
  h2 {
    font-size: var(--text-xl);
  }
  h3 {
    font-size: var(--text-lg);
  }
  a {
    color: var(--color-blue-600);
    text-decoration-line: underline;
  }
}
```

--------------------------------

### Applying Responsive Font Stretch Utilities in HTML

Source: https://tailwindcss.com/docs/font-stretch

Prefix `font-stretch` utilities with a breakpoint variant like `md:` to apply them conditionally based on screen size.

```html
<div class="font-stretch-normal md:font-stretch-expanded ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Vertical Alignment

Source: https://tailwindcss.com/docs/vertical-align

Apply vertical alignment utilities at specific breakpoints using responsive prefixes like `md:`.

```html
<span class="align-middle md:align-top ...">  <!-- ... --></span>
```

--------------------------------

### Basic backdrop brightness with predefined utilities

Source: https://tailwindcss.com/docs/backdrop-filter-brightness

Apply backdrop brightness filters using Tailwind utility classes like backdrop-brightness-50 and backdrop-brightness-150 to control element backdrop brightness.

```html
<div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-brightness-50 ..."></div></div><div class="bg-[url(/img/mountains.jpg)]">  <div class="bg-white/30 backdrop-brightness-150 ..."></div></div>
```

--------------------------------

### Flex Items with Container Scale Basis

Source: https://tailwindcss.com/docs/flex-basis

Sets the initial size of flex items using the container scale. Utilities like `basis-xs` and `basis-sm` are available.

```html
<div class="flex flex-row">  <div class="basis-3xs">01</div>  <div class="basis-2xs">02</div>  <div class="basis-xs">03</div>  <div class="basis-sm">04</div></div>
```

--------------------------------

### Responsive object-position

Source: https://tailwindcss.com/docs/object-position

Prefix an object-position utility with a breakpoint variant like md: to apply the utility only at medium screen sizes and above. This allows for responsive adjustments to content positioning.

```html
<img class="object-center md:object-top ..." src="/img/mountains.jpg" />
```

--------------------------------

### Allowing a flex item to grow using Tailwind CSS

Source: https://tailwindcss.com/docs/flex-grow

Use the `grow` utility to make a flex item expand and fill any available space within its flex container.

```html
<div class="flex ...">  <div class="size-14 flex-none ...">01</div>  <div class="size-14 grow ...">02</div>  <div class="size-14 flex-none ...">03</div></div>
```

--------------------------------

### Registering External Library Sources

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Use the @source directive to register paths to external libraries relative to the stylesheet. This is useful for scanning dependencies ignored by default.

```css
@import "tailwindcss";@source "../node_modules/@acmecorp/ui-lib";
```

--------------------------------

### Apply predefined blur filters to an element

Source: https://tailwindcss.com/docs/filter-blur

Use utility classes like blur-none, blur-sm, blur-lg, and blur-2xl to apply different levels of blur to an image.

```html
<img class="blur-none" src="/img/mountains.jpg" /><img class="blur-sm" src="/img/mountains.jpg" /><img class="blur-lg" src="/img/mountains.jpg" /><img class="blur-2xl" src="/img/mountains.jpg" />
```

--------------------------------

### Scrolling background with container using bg-local

Source: https://tailwindcss.com/docs/background-attachment

Use bg-local to scroll the background image with the container and the viewport.

```html
<div class="bg-[url(/img/mountains.jpg)] bg-local ...">  <!-- ... --></div>
```

--------------------------------

### Providing Default Values with --default()

Source: https://tailwindcss.com/docs/adding-custom-styles

Use `--default()` within `--value()` to specify a default value when the utility is used without an explicit value.

```css
@utility tab-* {
  tab-size: --value(integer, --default(4));
}
```

--------------------------------

### Responsive mask-origin

Source: https://tailwindcss.com/docs/mask-origin

Prefix a mask-origin utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above. This allows for responsive control over mask positioning.

```html
<div class="mask-origin-border md:mask-origin-padding ...">  <!-- ... --></div>
```

--------------------------------

### Set Custom Opacity with CSS Variable in Tailwind CSS

Source: https://tailwindcss.com/docs/opacity

Demonstrates using the `opacity-(<custom-property>)` syntax to apply opacity based on a CSS variable.

```html
<button class="opacity-(--my-opacity) ...">  <!-- ... --></button>
```

--------------------------------

### Select All Text on Click

Source: https://tailwindcss.com/docs/user-select

Use the `select-all` utility to automatically select all the text within an element when a user clicks on it. This is convenient for easily copying content.

```html
<div class="select-all ...">The quick brown fox jumps over the lazy dog.</div>
```

--------------------------------

### Responsive border-collapse with breakpoint variant

Source: https://tailwindcss.com/docs/border-collapse

Use the md: breakpoint prefix to apply border-collapse only at medium screen sizes and above, allowing different border behavior at smaller breakpoints.

```html
<table class="border-collapse md:border-separate ...">
  <!-- ... -->
</table>
```

--------------------------------

### Setting Inline Border Color to Amber Shades

Source: https://tailwindcss.com/docs/border-color

Demonstrates setting the inline border color using various shades of amber, from `amber-50` (lightest) to `amber-950` (darkest).

```html
<!-- Set border-inline-color to amber-50 -->
<div class="border-x-amber-50"></div>

<!-- Set border-inline-color to amber-100 -->
<div class="border-x-amber-100"></div>

<!-- Set border-inline-color to amber-200 -->
<div class="border-x-amber-200"></div>

<!-- Set border-inline-color to amber-300 -->
<div class="border-x-amber-300"></div>

<!-- Set border-inline-color to amber-400 -->
<div class="border-x-amber-400"></div>

<!-- Set border-inline-color to amber-500 -->
<div class="border-x-amber-500"></div>

<!-- Set border-inline-color to amber-600 -->
<div class="border-x-amber-600"></div>

<!-- Set border-inline-color to amber-700 -->
<div class="border-x-amber-700"></div>

<!-- Set border-inline-color to amber-800 -->
<div class="border-x-amber-800"></div>

<!-- Set border-inline-color to amber-900 -->
<div class="border-x-amber-900"></div>

<!-- Set border-inline-color to amber-950 -->
<div class="border-x-amber-950"></div>
```

--------------------------------

### Negative Scaling and Mirroring

Source: https://tailwindcss.com/docs/scale

Use negative scale utilities like `-scale-x-75` and `-scale-y-125` to mirror and scale down an element. These utilities apply scaling along the respective axis.

```html
<img class="-scale-x-75 ..." src="/img/mountains.jpg" /><img class="-scale-100 ..." src="/img/mountains.jpg" /><img class="-scale-y-125 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Customizing Transition Timing Functions

Source: https://tailwindcss.com/docs/transition-timing-function

Customize the default transition timing function utilities by defining `--ease-*` theme variables within your `@theme` block. This allows you to introduce new easing curves like `ease-in-expo`.

```css
@theme {  --ease-in-expo: cubic-bezier(0.95, 0.05, 0.795, 0.035); }
```

```html
<button class="ease-in-expo">  <!-- ... --></button>
```

--------------------------------

### Customize Text Shadow Utility

Source: https://tailwindcss.com/docs/text-shadow

Customize the `--text-shadow-xl` theme variable to change the default text shadow utility. This allows for custom shadow effects in your project.

```css
@theme {  --text-shadow-xl: 0 35px 35px rgb(0, 0, 0 / 0.25); }
```

--------------------------------

### Responsive Stroke Width

Source: https://tailwindcss.com/docs/stroke-width

Prefix a `stroke-width` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="stroke-1 md:stroke-2 ...">  <!-- ... --></div>
```

--------------------------------

### Set Border Color with Stone Palette

Source: https://tailwindcss.com/docs/border-color

Utilize these classes to set border colors from the stone color palette. Each class applies a `border-inline-color` with a specific oklch value for the stone scale.

```html
`border-x-stone-50`| `border-inline-color: var(--color-stone-50); /* oklch(98.5% 0.001 106.423) */`
```

```html
`border-x-stone-100`| `border-inline-color: var(--color-stone-100); /* oklch(97% 0.001 106.424) */`
```

```html
`border-x-stone-200`| `border-inline-color: var(--color-stone-200); /* oklch(92.3% 0.003 48.717) */`
```

```html
`border-x-stone-300`| `border-inline-color: var(--color-stone-300); /* oklch(86.9% 0.005 56.366) */`
```

```html
`border-x-stone-400`| `border-inline-color: var(--color-stone-400); /* oklch(70.9% 0.01 56.259) */`
```

```html
`border-x-stone-500`| `border-inline-color: var(--color-stone-500); /* oklch(55.3% 0.013 58.071) */`
```

```html
`border-x-stone-600`| `border-inline-color: var(--color-stone-600); /* oklch(44.4% 0.011 73.639) */`
```

```html
`border-x-stone-700`| `border-inline-color: var(--color-stone-700); /* oklch(37.4% 0.01 67.558) */`
```

```html
`border-x-stone-800`| `border-inline-color: var(--color-stone-800); /* oklch(26.8% 0.007 34.298) */`
```

```html
`border-x-stone-900`| `border-inline-color: var(--color-stone-900); /* oklch(21.6% 0.006 56.043) */`
```

```html
`border-x-stone-950`| `border-inline-color: var(--color-stone-950); /* oklch(14.7% 0.004 49.25) */`
```

--------------------------------

### Transition-delay with CSS variables

Source: https://tailwindcss.com/docs/transition-delay

Use the `delay-(<custom-property>)` syntax as a shorthand for `delay-[var(<custom-property>)]` to set transition delay using CSS variables. The `var()` function is automatically added.

```html
<button class="delay-(--my-delay) ...">  <!-- ... --></button>
```

--------------------------------

### Basic brightness filter classes

Source: https://tailwindcss.com/docs/filter-brightness

Apply standard brightness levels to images using predefined utility classes. Classes range from brightness-50 to brightness-200 with corresponding percentage values.

```html
<img class="brightness-50 ..." src="/img/mountains.jpg" /><img class="brightness-100 ..." src="/img/mountains.jpg" /><img class="brightness-125 ..." src="/img/mountains.jpg" /><img class="brightness-200 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Arbitrary Spacing with --spacing() and calc()

Source: https://tailwindcss.com/docs/functions-and-directives

Combine --spacing() with calc() for arbitrary spacing values in HTML attributes. This allows for dynamic spacing calculations.

```html
<div class="py-[calc(--spacing(4)-1px)]">
  <!-- ... -->
</div>
```

--------------------------------

### Max-Width Container Queries

Source: https://tailwindcss.com/docs/responsive-design

Apply styles below a specific container size using max-width container query variants like `@max-sm` and `@max-md`.

```html
<div class="@container">  <div class="flex flex-row @max-md:flex-col">    <!-- ... -->  </div></div>
```

--------------------------------

### Apply GPU Hardware Acceleration for Transforms

Source: https://tailwindcss.com/docs/transform

Use `transform-gpu` to force hardware acceleration for transitions when performance is critical. This can be conditionally reverted with `transform-cpu`.

```html
<div class="scale-150 transform-gpu">  <!-- ... --></div>
```

--------------------------------

### Background Size: Custom Value

Source: https://tailwindcss.com/docs/background-size

Use `bg-size-[<value>]` for arbitrary background size values. This allows for precise control, such as setting height and width independently.

```html
<div class="bg-size-[auto_100px] ...">  <!-- ... --></div>
```

--------------------------------

### Dynamic CSS Variables with Utility Classes

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Set CSS variables using inline styles based on dynamic data, then reference these variables with Tailwind utility classes for dynamic theming.

```jsx
export function BrandedButton({ buttonColor, buttonColorHover, textColor, children }) {  return (    <button      style={{        "--bg-color": buttonColor,
        "--bg-color-hover": buttonColorHover,
        "--text-color": textColor,
      }}
      className="bg-(--bg-color) text-(--text-color) hover:bg-(--bg-color-hover) ..."
    >
      {children}
    </button>  );
}
```

--------------------------------

### Custom Sepia Value

Source: https://tailwindcss.com/docs/filter-sepia

Set a custom sepia amount using the `sepia-[<value>]` syntax for precise control.

```html
<img class="sepia-[.25] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Import app.css into app.js

Source: https://tailwindcss.com/docs/installation/framework-guides/emberjs

Import the main application CSS file into `app.js` to ensure it's loaded by Ember.

```javascript
import Application from '@ember/application';import Resolver from 'ember-resolver';import loadInitializers from 'ember-load-initializers';import config from 'my-project/config/environment';import 'my-project/app.css';export default class App extends Application {  modulePrefix = config.modulePrefix;  podModulePrefix = config.podModulePrefix;  Resolver = Resolver;}loadInitializers(App, config.modulePrefix);
```

--------------------------------

### place-self-stretch alignment in grid

Source: https://tailwindcss.com/docs/place-self

Stretch a grid item on both axes using the place-self-stretch class.

```html
<div class="grid grid-cols-3 gap-4 ...">
  <div>01</div>
  <div class="place-self-stretch ...">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### Apply a custom blur value using arbitrary values

Source: https://tailwindcss.com/docs/filter-blur

Set a precise blur amount using the blur-[<value>] syntax with any valid CSS length unit.

```html
<img class="blur-[2px] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Create completely custom theme in Tailwind CSS

Source: https://tailwindcss.com/docs/theme

Disable the default theme entirely by setting --*: initial, then define only your custom theme variables. Only utilities matching your custom variables will be available.

```CSS
@import "tailwindcss";
@theme {
  --*: initial;
  --spacing: 4px;
  --font-body: Inter, sans-serif;
  --color-lagoon: oklch(0.72 0.11 221.19);
  --color-coral: oklch(0.74 0.17 40.24);
  --color-driftwood: oklch(0.79 0.06 74.59);
  --color-tide: oklch(0.49 0.08 205.88);
  --color-dusk: oklch(0.82 0.15 72.09);
}
```

--------------------------------

### place-self-center alignment in grid

Source: https://tailwindcss.com/docs/place-self

Center a grid item on both axes using the place-self-center class.

```html
<div class="grid grid-cols-3 gap-4 ...">
  <div>01</div>
  <div class="place-self-center ...">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### Start-end logical border radius with preset sizes

Source: https://tailwindcss.com/docs/border-radius

Apply border-radius to the start-end corner using logical properties. Supports sizes from xs (0.125rem) to 4xl (2rem).

```css
border-start-end-radius: var(--radius-xs); /* 0.125rem (2px) */
```

```css
border-start-end-radius: var(--radius-sm); /* 0.25rem (4px) */
```

```css
border-start-end-radius: var(--radius-md); /* 0.375rem (6px) */
```

```css
border-start-end-radius: var(--radius-lg); /* 0.5rem (8px) */
```

```css
border-start-end-radius: var(--radius-xl); /* 0.75rem (12px) */
```

```css
border-start-end-radius: var(--radius-2xl); /* 1rem (16px) */
```

```css
border-start-end-radius: var(--radius-3xl); /* 1.5rem (24px) */
```

```css
border-start-end-radius: var(--radius-4xl); /* 2rem (32px) */
```

--------------------------------

### Define shared theme variables in separate CSS file

Source: https://tailwindcss.com/docs/theme

Create a reusable theme file with custom variables that can be imported across multiple projects in a monorepo or published to NPM.

```CSS
@theme {
  --*: initial;
  --spacing: 4px;
  --font-body: Inter, sans-serif;
  --color-lagoon: oklch(0.72 0.11 221.19);
  --color-coral: oklch(0.74 0.17 40.24);
  --color-driftwood: oklch(0.79 0.06 74.59);
  --color-tide: oklch(0.49 0.08 205.88);
  --color-dusk: oklch(0.82 0.15 72.09);
}
```

--------------------------------

### Responsive Text Decoration Color

Source: https://tailwindcss.com/docs/text-decoration-color

Apply text decoration color utilities at specific breakpoints using prefix variants like `md:`.

```html
<p class="underline decoration-sky-600 md:decoration-blue-400 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Set Columns by Number

Source: https://tailwindcss.com/docs/columns

Use `columns-<number>` utilities to specify the exact number of columns for content. The column width adjusts automatically.

```html
<div class="columns-3 ...">
  <img class="aspect-3/2 ..." src="/img/mountains-1.jpg" />
  <img class="aspect-square ..." src="/img/mountains-2.jpg" />
  <img class="aspect-square ..." src="/img/mountains-3.jpg" />
  <!-- ... -->
</div>
```

--------------------------------

### Apply Custom Saturation Value with Tailwind CSS

Source: https://tailwindcss.com/docs/filter-saturate

Set a custom saturation level using the `saturate-[<value>]` syntax, allowing for precise control beyond predefined utility classes.

```html
<img class="saturate-[.25] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Distribute content with space between in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-between` to distribute grid items with equal space between rows and columns along both axes.

```html
<div class="grid h-48 grid-cols-2 place-content-between gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Responsive text-decoration-style with breakpoint variant

Source: https://tailwindcss.com/docs/text-decoration-style

Use the md: breakpoint prefix to apply decoration-dashed only at medium screen sizes and above. Responsive variants allow conditional application of utilities based on screen width.

```html
<p class="underline md:decoration-dashed ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Apply padding to all sides with p-<number>

Source: https://tailwindcss.com/docs/padding

Use `p-<number>` utilities to control the padding on all sides of an element.

```html
<div class="p-8 ...">p-8</div>
```

--------------------------------

### Setting Font Stretch with Percentage Utilities in HTML

Source: https://tailwindcss.com/docs/font-stretch

Apply `font-stretch-<percentage>` utilities to set the font width using a specific percentage value.

```html
<p class="font-stretch-50%">The quick brown fox...</p><p class="font-stretch-100%">The quick brown fox...</p><p class="font-stretch-150%">The quick brown fox...</p>
```

--------------------------------

### Load Legacy JavaScript Plugin with @plugin

Source: https://tailwindcss.com/docs/functions-and-directives

Use the @plugin directive to load a legacy JavaScript-based plugin. It accepts either a package name or a local path.

```css
@plugin "@tailwindcss/typography";
```

--------------------------------

### Style Lists with Utilities

Source: https://tailwindcss.com/docs/preflight

Lists can be styled using Tailwind's `list-style-type` and `list-style-position` utilities.

```html
<ul class="list-inside list-disc">
  <li>One</li>
  <li>Two</li>
  <li>Three</li>
</ul>
```

--------------------------------

### Set border-inline-color with arbitrary values

Source: https://tailwindcss.com/docs/border-color

Use the `border-x-` utility with square brackets to set an arbitrary inline border color value.

```html
<div class="border-x-[red]">...</div>
```

--------------------------------

### Predefined Box Shadows

Source: https://tailwindcss.com/docs/box-shadow

Apply predefined box shadow effects to elements using utility classes. These range from subtle to pronounced shadows.

```html
<div class="shadow-2xs"></div>
<div class="shadow-xs"></div>
<div class="shadow-sm"></div>
<div class="shadow-md"></div>
<div class="shadow-lg"></div>
<div class="shadow-xl"></div>
<div class="shadow-2xl"></div>
<div class="shadow-none"></div>
```

--------------------------------

### Using CSS custom properties for text-decoration-color

Source: https://tailwindcss.com/docs/text-decoration-color

Apply text decoration colors using CSS custom properties by utilizing the `decoration-[<custom-property>]` utility. This allows for dynamic color control through variables.

```html
<p class="decoration-[--color-my-custom-color]">Custom Color</p>
```

--------------------------------

### Load Legacy JavaScript Configuration with @config

Source: https://tailwindcss.com/docs/functions-and-directives

Use the @config directive to load a legacy JavaScript-based configuration file. Note that corePlugins, safelist, and separator options are not supported in v4.0.

```css
@config "../../tailwind.config.js";
```

--------------------------------

### Import Tailwind with 'prefix' option

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Use the 'prefix' option when importing Tailwind CSS to prefix all generated classes and CSS variables. This helps avoid conflicts with existing class names.

```css
@import "tailwindcss" prefix(tw);
```

--------------------------------

### Display element content at original size with object-none (HTML)

Source: https://tailwindcss.com/docs/object-fit

Use the `object-none` utility to display an element's content at its original size ignoring the container size.

```html
<img class="h-48 w-96 object-none ..." src="/img/mountains.jpg" />
```

--------------------------------

### Ignore Pointer Events with pointer-events-none

Source: https://tailwindcss.com/docs/pointer-events

Use the `pointer-events-none` utility to make an element and its children ignore pointer events. Events will pass through to elements beneath.

```html
<div class="relative ...">  <div class="pointer-events-auto absolute ...">    <svg class="absolute h-5 w-5 text-gray-400">      <!-- ... -->    </svg>  </div>  <input type="text" placeholder="Search" class="..." /></div><div class="relative ...">  <div class="pointer-events-none absolute ...">    <svg class="absolute h-5 w-5 text-gray-400">      <!-- ... -->    </svg>  </div>  <input type="text" placeholder="Search" class="..." /></div>
```

--------------------------------

### Backdrop Sepia with CSS Variables

Source: https://tailwindcss.com/docs/backdrop-filter-sepia

For CSS variables, use the backdrop-sepia-(<custom-property>) syntax. This is a shorthand for backdrop-sepia-[var(<custom-property>)].

```html
<div class="backdrop-sepia-(--my-backdrop-sepia) ...">  <!-- ... --></div>
```

--------------------------------

### Use Custom Column Value

Source: https://tailwindcss.com/docs/columns

Employ the `columns-[<value>]` syntax for arbitrary column values, such as `30vw`. For CSS variables, use `columns-(<custom-property>)`.

```html
<div class="columns-[30vw] ...">
  <!-- ... -->
</div>
```

```html
<div class="columns-(--my-columns) ...">
  <!-- ... -->
</div>
```

--------------------------------

### Using diagonal fractions

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `diagonal-fractions` utility to replace numbers separated by a slash with common diagonal fractions in fonts that support them.

```html
<p class="diagonal-fractions ...">1/2 3/4 5/6</p>
```

--------------------------------

### Setting Font Size and Line Height Together

Source: https://tailwindcss.com/docs/line-height

Use utilities like `text-sm/6` to set both the font size and line height simultaneously. This is useful for maintaining consistent vertical rhythm based on font size.

```html
<p class="text-base/6 ...">So I started to walk into the water...</p><p class="text-base/7 ...">So I started to walk into the water...</p><p class="text-base/8 ...">So I started to walk into the water...</p>
```

--------------------------------

### Globally add CSS file to Nuxt configuration

Source: https://tailwindcss.com/docs/installation/framework-guides/nuxt

Reference the `main.css` file in the `css` array of `nuxt.config.ts` to ensure global styling.

```typescript
import tailwindcss from "@tailwindcss/vite";export default defineNuxtConfig({  compatibilityDate: "2025-07-15",  devtools: { enabled: true },  css: ['./app/assets/css/main.css'],  vite: {    plugins: [      tailwindcss(),    ],  },});
```

--------------------------------

### Apply Responsive Font Weight in HTML

Source: https://tailwindcss.com/docs/font-weight

Conditionally apply font weight utilities at different screen sizes by prefixing them with breakpoint variants like `md:`.

```html
<p class="font-normal md:font-bold ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Transition with custom CSS variable property

Source: https://tailwindcss.com/docs/transition-property

Use the `transition-(<custom-property>)` syntax as a shorthand for `transition-[var(<custom-property>)]` to transition custom CSS variables. This automatically adds the `var()` function.

```html
<button class="transition-(--my-properties) ...">  <!-- ... --></button>
```

--------------------------------

### Custom Drop Shadow with CSS Variables

Source: https://tailwindcss.com/docs/filter-drop-shadow

Use the `drop-shadow-(<custom-property>)` syntax as a shorthand for applying drop shadows defined by CSS variables.

```html
<svg class="drop-shadow-(--my-drop-shadow) ...">  <!-- ... --></svg>
```

--------------------------------

### Apply Responsive Grid Row Spanning

Source: https://tailwindcss.com/docs/grid-row

Prefix grid row utilities with breakpoint variants like `md:` to apply them conditionally based on screen size. This enables responsive grid layouts.

```html
<div class="row-span-3 md:row-span-4 ...">  <!-- ... --></div>
```

--------------------------------

### Handling Whitespace in Arbitrary Values

Source: https://tailwindcss.com/docs/adding-custom-styles

Use an underscore (_) instead of a space in arbitrary values to ensure correct conversion to spaces at build-time.

```html
<div class="grid grid-cols-[1fr_500px_2fr]">  <!-- ... --></div>
```

--------------------------------

### Center content in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-center` to horizontally and vertically center grid items within their container.

```html
<div class="grid h-48 grid-cols-2 place-content-center gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### End-start logical border radius with preset sizes

Source: https://tailwindcss.com/docs/border-radius

Apply border-radius to the end-start corner using logical properties. Supports sizes from xs (0.125rem) to 4xl (2rem).

```css
border-end-start-radius: var(--radius-xs); /* 0.125rem (2px) */
```

```css
border-end-start-radius: var(--radius-sm); /* 0.25rem (4px) */
```

```css
border-end-start-radius: var(--radius-md); /* 0.375rem (6px) */
```

```css
border-end-start-radius: var(--radius-lg); /* 0.5rem (8px) */
```

```css
border-end-start-radius: var(--radius-xl); /* 0.75rem (12px) */
```

```css
border-end-start-radius: var(--radius-2xl); /* 1rem (16px) */
```

```css
border-end-start-radius: var(--radius-3xl); /* 1.5rem (24px) */
```

```css
border-end-start-radius: var(--radius-4xl); /* 2rem (32px) */
```

--------------------------------

### Responsive accent color design

Source: https://tailwindcss.com/docs/accent-color

Prefix an `accent-color` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<input class="accent-black md:accent-pink-500 ..." type="checkbox" />
```

--------------------------------

### Applying an Image Mask

Source: https://tailwindcss.com/docs/mask-image

Use the `mask-[<value>]` utility to set a custom image as a mask for an element. Ensure the background image is also defined.

```html
<div class="mask-[url(/img/scribble.png)] bg-[url(/img/mountains.jpg)] ...">  <!-- ... --></div>
```

--------------------------------

### Set Background Image

Source: https://tailwindcss.com/docs/background-image

Use the `bg-[<value>]` syntax to set a custom background image for an element. Ensure the image path is correctly specified.

```html
<div class="bg-[url(/img/mountains.jpg)] ..."></div>
```

--------------------------------

### Using tabular figures

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `tabular-nums` utility to use numeric glyphs that have uniform/tabular widths in fonts that support them.

```html
<p class="tabular-nums ...">12121</p><p class="tabular-nums ...">90909</p>
```

--------------------------------

### Predefined text-decoration-color utilities

Source: https://tailwindcss.com/docs/text-decoration-color

Use these classes to apply predefined colors to text decorations. Each class maps to a specific CSS `text-decoration-color` property value, often using CSS variables for theming.

```html
<p class="decoration-mist-900">Mist 900</p>
<p class="decoration-mist-950">Mist 950</p>
<p class="decoration-taupe-50">Taupe 50</p>
<p class="decoration-taupe-100">Taupe 100</p>
<p class="decoration-taupe-200">Taupe 200</p>
<p class="decoration-taupe-300">Taupe 300</p>
<p class="decoration-taupe-400">Taupe 400</p>
<p class="decoration-taupe-500">Taupe 500</p>
<p class="decoration-taupe-600">Taupe 600</p>
<p class="decoration-taupe-700">Taupe 700</p>
<p class="decoration-taupe-800">Taupe 800</p>
<p class="decoration-taupe-900">Taupe 900</p>
<p class="decoration-taupe-950">Taupe 950</p>
```

--------------------------------

### Pretty Text Wrapping

Source: https://tailwindcss.com/docs/text-wrap

Use the `text-pretty` utility to prefer better text wrapping and layout, potentially preventing orphans at the expense of speed. Browser behavior may vary.

```html
<article>  <h3 class="text-pretty">Beloved Manhattan soup stand closes</h3>  <p>New Yorkers are facing the winter chill...</p></article>
```

--------------------------------

### Sepia with CSS Variable

Source: https://tailwindcss.com/docs/filter-sepia

Apply sepia using a CSS variable with the `sepia-(<custom-property>)` syntax, which is a shorthand for `sepia-[var(<custom-property>)]`.

```html
<img class="sepia-(--my-sepia) ..." src="/img/mountains.jpg" />
```

--------------------------------

### Set Gradient Stop Positions

Source: https://tailwindcss.com/docs/background-image

Precisely control the positions of gradient color stops using percentage values with `from-<percentage>`, `via-<percentage>`, and `to-<percentage>` utilities. This allows for fine-tuning the gradient appearance.

```html
<div class="bg-linear-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90% ..."></div>
```

--------------------------------

### Basic Outline Width

Source: https://tailwindcss.com/docs/outline-width

Set the width of an element's outline using `outline` or `outline-<number>` utilities like `outline-2` and `outline-4`. These can be combined with `outline-offset`.

```html
<button class="outline outline-offset-2 ...">Button A</button><button class="outline-2 outline-offset-2 ...">Button B</button><button class="outline-4 outline-offset-2 ...">Button C</button>
```

--------------------------------

### Using Negative Positioning Values

Source: https://tailwindcss.com/docs/top-right-bottom-left

Prefix the positioning utility class with a dash to apply a negative value, moving the element outside its normal flow.

```html
<div class="relative size-32 ...">  <div class="absolute -top-4 -left-4 size-14 ..."></div></div>
```

--------------------------------

### Make Images Inline

Source: https://tailwindcss.com/docs/preflight

Use the `inline` utility to override the default `display: block` for images and replaced elements.

```html
<img class="inline" src="..." alt="..." />
```

--------------------------------

### Supporting Literal Values

Source: https://tailwindcss.com/docs/adding-custom-styles

Use the `--value('literal')` syntax to support specific literal string values for a utility.

```css
@utility tab-* {
  tab-size: --value("inherit", "initial", "unset");
}
```

--------------------------------

### Apply font-feature-settings using CSS Variables

Source: https://tailwindcss.com/docs/font-feature-settings

Apply font feature settings from a CSS variable using the `font-features-(<custom-property>)` syntax.

```html
<p class="font-features-(--my-features) ...">  <!-- ... --></p>
```

--------------------------------

### Flex Items with Spacing Scale Basis

Source: https://tailwindcss.com/docs/flex-basis

Sets the initial size of flex items using the spacing scale. Use utilities like `basis-64` and `basis-128`.

```html
<div class="flex flex-row">  <div class="basis-64">01</div>  <div class="basis-64">02</div>  <div class="basis-128">03</div></div>
```

--------------------------------

### Basic border radius sizes

Source: https://tailwindcss.com/docs/border-radius

Apply different border radius sizes using utilities like rounded-sm, rounded-md, rounded-lg, and rounded-xl.

```html
<div class="rounded-sm ..."></div><div class="rounded-md ..."></div><div class="rounded-lg ..."></div><div class="rounded-xl ..."></div>
```

--------------------------------

### Configure Vite plugin for Tailwind CSS

Source: https://tailwindcss.com/docs/installation/framework-guides/sveltekit

Adds the `@tailwindcss/vite` plugin to the `vite.config.ts` file to enable Tailwind CSS processing in your SvelteKit project.

```typescript
import { sveltekit } from '@sveltejs/kit/vite';import { defineConfig } from 'vite';import tailwindcss from '@tailwindcss/vite';export default defineConfig({
  plugins: [
    tailwindcss(),
    sveltekit(),
  ],
});
```

--------------------------------

### Custom Tab Size with CSS Variables

Source: https://tailwindcss.com/docs/tab-size

For CSS variables, use the `tab-(<custom-property>)` syntax, which is a shorthand for `tab-[var(<custom-property>)]`.

```html
<pre class="tab-(--my-tab-size) ...">  <!-- ... --></pre>
```

--------------------------------

### Basic font-family utilities

Source: https://tailwindcss.com/docs/font-family

Use font-sans, font-serif, and font-mono to set the font family for an element.

```html
<p class="font-sans ...">The quick brown fox ...</p><p class="font-serif ...">The quick brown fox ...</p><p class="font-mono ...">The quick brown fox ...</p>
```

--------------------------------

### Set percentage-based maximum height

Source: https://tailwindcss.com/docs/max-height

Use max-h-full or max-h-<fraction> utilities to set a maximum height as a percentage of the parent's height.

```html
<div class="h-96 ...">  <div class="h-full max-h-9/10 ...">max-h-9/10</div>  <div class="h-full max-h-3/4 ...">max-h-3/4</div>  <div class="h-full max-h-1/2 ...">max-h-1/2</div>  <div class="h-full max-h-1/4 ...">max-h-1/4</div>  <div class="h-full max-h-full ...">max-h-full</div></div>
```

--------------------------------

### Basic Inline Sizes

Source: https://tailwindcss.com/docs/inline-size

Use `inline-<number>` utilities to set an element to a fixed inline size based on the spacing scale.

```html
<div class="inline-96 ...">inline-96</div><div class="inline-80 ...">inline-80</div><div class="inline-64 ...">inline-64</div><div class="inline-48 ...">inline-48</div><div class="inline-40 ...">inline-40</div><div class="inline-32 ...">inline-32</div>
```

--------------------------------

### Custom Backdrop Contrast Value

Source: https://tailwindcss.com/docs/backdrop-filter-contrast

Use the `backdrop-contrast-[<value>]` syntax to set the backdrop contrast based on a completely custom value. This applies `backdrop-filter: contrast(<value>);`.

```html
<div class="backdrop-contrast-[.25] ...">  <!-- ... --></div>
```

--------------------------------

### Customizing Aspect Ratio Theme

Source: https://tailwindcss.com/docs/aspect-ratio

Use the `--aspect-*` theme variables to customize the aspect ratio utilities in your project. The `aspect-retro` utility can then be used in your markup.

```css
@theme {  --aspect-retro: 4 / 3; }
```

```html
<iframe class="aspect-retro" src="https://www.youtube.com/embed/dQw4w9WgXcQ"></iframe>
```

--------------------------------

### Predefined Background Colors (Olive Scale)

Source: https://tailwindcss.com/docs/background-color

Applies background colors from the olive color scale. Each class maps to a specific CSS variable.

```html
<div class="bg-olive-700">...</div>
<div class="bg-olive-800">...</div>
<div class="bg-olive-900">...</div>
<div class="bg-olive-950">...</div>
```

--------------------------------

### Responsive Backdrop Hue Rotate

Source: https://tailwindcss.com/docs/backdrop-filter-hue-rotate

Prefix a `backdrop-hue-rotate()` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above. This enables responsive adjustments to backdrop hue rotation.

```html
<div class="backdrop-hue-rotate-15 md:backdrop-hue-rotate-0 ...">  <!-- ... --></div>
```

--------------------------------

### Custom Value max-inline-size Utilities

Source: https://tailwindcss.com/docs/max-inline-size

Apply custom maximum inline sizes using the `max-inline-[<value>]` syntax for arbitrary pixel values or CSS variables. This provides maximum flexibility for specific design requirements.

```html
<div class="max-inline-[220px] ...">  <!-- ... --></div>
```

```html
<div class="max-inline-(--my-max-inline-size) ...">  <!-- ... --></div>
```

--------------------------------

### Custom Background Color with CSS Variables

Source: https://tailwindcss.com/docs/background-color

Apply background colors using CSS variables by using the `bg-(<custom-property>)` syntax, which is a shorthand for `bg-[var(<custom-property>)]`.

```html
<div class="bg-(--my-color) ...">  <!-- ... --></div>
```

--------------------------------

### Basic Transition Easing

Source: https://tailwindcss.com/docs/transition-timing-function

Use `ease-in`, `ease-out`, and `ease-in-out` to control the easing curve of an element's transition. Apply these classes to elements to observe their behavior on hover.

```html
<button class="duration-300 ease-in ...">Button A</button><button class="duration-300 ease-out ...">Button B</button><button class="duration-300 ease-in-out ...">Button C</button>
```

--------------------------------

### Combining Theme, Bare, and Arbitrary Values

Source: https://tailwindcss.com/docs/adding-custom-styles

Combine multiple `--value()` declarations within a rule to support theme, bare, and arbitrary values simultaneously. Unresolvable declarations are omitted.

```css
@theme {
  --tab-size-github: 8;
}
@utility tab-* {
  tab-size: --value([integer]);
  tab-size: --value(integer);
  tab-size: --value(--tab-size-*);
}
```

--------------------------------

### Responsive Backdrop Grayscale

Source: https://tailwindcss.com/docs/backdrop-filter-grayscale

Prefix a `backdrop-filter: grayscale()` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="backdrop-grayscale md:backdrop-grayscale-0 ...">  <!-- ... --></div>
```

--------------------------------

### Predefined Background Colors (Taupe Scale)

Source: https://tailwindcss.com/docs/background-color

Applies background colors from the taupe color scale. Each class maps to a specific CSS variable.

```html
<div class="bg-taupe-50">...</div>
<div class="bg-taupe-100">...</div>
<div class="bg-taupe-200">...</div>
<div class="bg-taupe-300">...</div>
<div class="bg-taupe-400">...</div>
<div class="bg-taupe-500">...</div>
<div class="bg-taupe-600">...</div>
<div class="bg-taupe-700">...</div>
<div class="bg-taupe-800">...</div>
<div class="bg-taupe-900">...</div>
<div class="bg-taupe-950">...</div>
```

--------------------------------

### Applying Basic Font Stretch Utilities in HTML

Source: https://tailwindcss.com/docs/font-stretch

Use predefined `font-stretch` utilities to adjust the width of a font face. This applies only if the font has multiple width variations available.

```html
<p class="font-stretch-extra-condensed">The quick brown fox...</p><p class="font-stretch-condensed">The quick brown fox...</p><p class="font-stretch-normal">The quick brown fox...</p><p class="font-stretch-expanded">The quick brown fox...</p><p class="font-stretch-extra-expanded">The quick brown fox...</p>
```

--------------------------------

### Responsive background-origin

Source: https://tailwindcss.com/docs/background-origin

Prefix a `background-origin` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above. This allows for responsive adjustments to background positioning.

```html
<div class="bg-origin-border md:bg-origin-padding ...">  <!-- ... --></div>
```

--------------------------------

### Set Outline Color

Source: https://tailwindcss.com/docs/outline-color

Use utilities like `outline-rose-500` and `outline-lime-100` to control the color of an element's outline. Ensure the `outline-width` utility is also applied.

```html
<button class="outline-2 outline-offset-2 outline-blue-500 ...">Button A</button><button class="outline-2 outline-offset-2 outline-cyan-500 ...">Button B</button><button class="outline-2 outline-offset-2 outline-pink-500 ...">Button C</button>
```

--------------------------------

### Adjusting Tailwind CSS Color Opacity

Source: https://tailwindcss.com/docs/colors

Demonstrates adjusting the opacity of a background color using the `/` syntax in Tailwind CSS, showcasing a range of alpha values.

```html
<div>  <div class="bg-sky-500/10"></div>  <div class="bg-sky-500/20"></div>  <div class="bg-sky-500/30"></div>  <div class="bg-sky-500/40"></div>  <div class="bg-sky-500/50"></div>  <div class="bg-sky-500/60"></div>  <div class="bg-sky-500/70"></div>  <div class="bg-sky-500/80"></div>  <div class="bg-sky-500/90"></div>  <div class="bg-sky-500/100"></div></div>
```

--------------------------------

### Style Placeholder-Shown Inputs with :placeholder-shown Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles to an input element when its placeholder is currently visible.

```html
<input class="placeholder-shown:border-gray-500 ..." placeholder="you@example.com" />
```

--------------------------------

### Arbitrary Values for Styling Properties

Source: https://tailwindcss.com/docs/adding-custom-styles

Apply arbitrary values to various CSS properties like background colors, font sizes, and pseudo-elements.

```html
<div class="bg-[#bada55] text-[22px] before:content-['Festivus']">  <!-- ... --></div>
```

--------------------------------

### Set pseudo-element content with Tailwind CSS

Source: https://tailwindcss.com/docs/content

Use the `content-[<value>]` syntax with `before` and `after` variants to set the content for `::before` and `::after` pseudo-elements. This is useful for adding decorative text or icons.

```html
<p>Higher resolution means more than just a better-quality image. With aRetina 6K display, <a class="text-blue-600 after:content-['_↗']" href="...">Pro Display XDR</a> gives you nearly 40 percent more screen real estate thana 5K display.</p>
```

--------------------------------

### Basic Text Indent

Source: https://tailwindcss.com/docs/text-indent

Use `indent-<number>` utilities like `indent-2` and `indent-8` to set the amount of empty space (indentation) that's shown before text in a block.

```html
<p class="indent-8">So I started to walk into the water...</p>
```

--------------------------------

### Arbitrary Background Color

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Use square bracket syntax for one-off colors outside the theme palette.

```html
<button class="bg-[#316ff6] ...">  Sign in with Facebook</button>
```

--------------------------------

### Create pill buttons

Source: https://tailwindcss.com/docs/border-radius

Use the rounded-full utility to create fully rounded pill-shaped buttons.

```html
<button class="rounded-full ...">Save Changes</button>
```

--------------------------------

### Update Ring Utility Width

Source: https://tailwindcss.com/docs/upgrade-guide

Replace `ring` with `ring-3` to adjust for the default ring width change from 3px to 1px in v4.

```html
<input class="ring ring-blue-500" /><input class="ring-3 ring-blue-500" />
```

--------------------------------

### Named Container Queries

Source: https://tailwindcss.com/docs/responsive-design

For complex layouts, name containers using `@container/{name}` and target them with specific variants like `@sm/{name}`. This allows styling based on distant containers.

```html
<div class="@container/main">  <!-- ... -->  <div class="flex flex-row @sm/main:flex-col">    <!-- ... -->  </div></div>
```

--------------------------------

### Responsive Mask Positioning

Source: https://tailwindcss.com/docs/mask-position

Apply mask position utilities responsively by prefixing them with a breakpoint variant like `md:`. This ensures the utility is applied only at the specified screen size and above.

```html
<div class="mask-center md:mask-top ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Border Color

Source: https://tailwindcss.com/docs/border-color

Prefix a `border-color` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<div class="border-blue-500 md:border-green-500 ...">  <!-- ... --></div>
```

--------------------------------

### Set Margin Utilities in Tailwind CSS

Source: https://tailwindcss.com/docs/margin

Use these utilities to apply consistent margin values based on the spacing scale. Negative values and `auto` are supported.

```html
<div class="m-4">...</div>
<div class="-m-4">...</div>
<div class="m-auto">...</div>
<div class="m-px">...</div>
<div class="-m-px">...</div>
```

```html
<div class="mx-4">...</div>
<div class="-mx-4">...</div>
<div class="mx-auto">...</div>
<div class="mx-px">...</div>
<div class="-mx-px">...</div>
```

```html
<div class="my-4">...</div>
<div class="-my-4">...</div>
<div class="my-auto">...</div>
<div class="my-px">...</div>
<div class="-my-px">...</div>
```

```html
<div class="ms-4">...</div>
<div class="-ms-4">...</div>
<div class="ms-auto">...</div>
<div class="ms-px">...</div>
<div class="-ms-px">...</div>
```

```html
<div class="me-4">...</div>
<div class="-me-4">...</div>
<div class="me-auto">...</div>
<div class="me-px">...</div>
<div class="-me-px">...</div>
```

```html
<div class="mbs-4">...</div>
<div class="-mbs-4">...</div>
<div class="mbs-auto">...</div>
<div class="mbs-px">...</div>
<div class="-mbs-px">...</div>
```

```html
<div class="mbe-4">...</div>
<div class="-mbe-4">...</div>
<div class="mbe-auto">...</div>
<div class="mbe-px">...</div>
<div class="-mbe-px">...</div>
```

```html
<div class="mt-4">...</div>
<div class="-mt-4">...</div>
<div class="mt-auto">...</div>
<div class="mt-px">...</div>
<div class="-mt-px">...</div>
```

```html
<div class="mr-4">...</div>
<div class="-mr-4">...</div>
<div class="mr-auto">...</div>
<div class="mr-px">...</div>
<div class="-mr-px">...</div>
```

```html
<div class="mb-4">...</div>
<div class="-mb-4">...</div>
<div class="mb-auto">...</div>
<div class="mb-px">...</div>
<div class="-mb-px">...</div>
```

```html
<div class="ml-4">...</div>
<div class="-ml-4">...</div>
<div class="ml-auto">...</div>
<div class="ml-px">...</div>
<div class="-ml-px">...</div>
```

--------------------------------

### Flex item shrink only, respecting initial size

Source: https://tailwindcss.com/docs/flex

Use `flex-initial` to allow a flex item to shrink but not grow, taking into account its initial size.

```html
<div class="flex">  <div class="w-14 flex-none ...">01</div>  <div class="w-64 flex-initial ...">02</div>  <div class="w-32 flex-initial ...">03</div></div>
```

--------------------------------

### Match Small Viewport Block Size

Source: https://tailwindcss.com/docs/block-size

Use the `block-svh` utility to set an element's block size to the smallest possible size of the viewport.

```html
<div class="block-svh">  <!-- ... --></div>
```

--------------------------------

### Responsive Letter Spacing with Breakpoints

Source: https://tailwindcss.com/docs/letter-spacing

Apply different letter spacing based on screen size using responsive prefixes like `md:tracking-wide` to adjust typography for various viewports.

```html
<p class="tracking-tight md:tracking-wide ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Inset shadow color utilities

Source: https://tailwindcss.com/docs/box-shadow

Control the color of inset shadows using utility classes. These classes set the `--tw-inset-shadow-color` custom property.

```html
<div class="inset-shadow-inherit">...</div>
<div class="inset-shadow-current">...</div>
<div class="inset-shadow-transparent">...</div>
```

--------------------------------

### Distribute content evenly in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-evenly` to distribute grid items with even spacing along both the inline and block axes.

```html
<div class="grid h-48 grid-cols-2 place-content-evenly gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Supporting Arbitrary Values with Data Types

Source: https://tailwindcss.com/docs/adding-custom-styles

Use the `--value([{type}])` syntax to allow arbitrary values of specified data types for a utility.

```css
@utility tab-* {
  tab-size: --value([integer]);
}
```

--------------------------------

### Responsive mask-repeat design

Source: https://tailwindcss.com/docs/mask-repeat

Prefix a `mask-repeat` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above. This allows for adaptive mask repetition.

```html
<div class="mask-repeat md:mask-repeat-x ...">  <!-- ... --></div>
```

--------------------------------

### Whitespace Pre Line Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-pre-line` utility to preserve newlines but not spaces. Text will be wrapped normally.

```html
<p class="whitespace-pre-line">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Flow Root for Block Formatting Context

Source: https://tailwindcss.com/docs/display

Apply the 'flow-root' utility to create a new block formatting context, useful for containing floats or managing element flow.

```html
<div class="p-4">  <div class="flow-root ...">    <div class="my-4 ...">Well, let me tell you something, ...</div>  </div>  <div class="flow-root ...">    <div class="my-4 ...">Sure, go ahead, laugh if you want...</div>  </div></div>
```

--------------------------------

### Start-end logical border radius with arbitrary value

Source: https://tailwindcss.com/docs/border-radius

Apply start-end border-radius using an arbitrary CSS value.

```css
border-start-end-radius: <value>;
```

--------------------------------

### Responsive Text Indent

Source: https://tailwindcss.com/docs/text-indent

Prefix a `text-indent` utility with a breakpoint variant like `md:` to only apply the utility at medium screen sizes and above.

```html
<p class="indent-4 md:indent-8 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Arbitrary Properties with Modifiers

Source: https://tailwindcss.com/docs/adding-custom-styles

Apply arbitrary CSS properties and combine them with modifiers like 'hover' for dynamic styling.

```html
<div class="[mask-type:luminance] hover:[mask-type:alpha]">  <!-- ... --></div>
```

--------------------------------

### Constrain Images Responsively

Source: https://tailwindcss.com/docs/preflight

Images and videos are constrained to their parent width with `max-width: 100%` and `height: auto` for responsiveness.

```css
img,video {
  max-width: 100%;
  height: auto;
}
```

--------------------------------

### Customize Inset Box Shadow - Tailwind CSS

Source: https://tailwindcss.com/docs/box-shadow

Use the --inset-shadow-* theme variables to customize the inset box shadow utilities. This enables the creation of custom inset shadows for elements.

```css
@theme {  --inset-shadow-md: inset 0 2px 3px rgba(0, 0, 0, 0.25); }
```

```html
<div class="inset-shadow-md">  <!-- ... --></div>
```

--------------------------------

### Custom text-decoration-thickness value

Source: https://tailwindcss.com/docs/text-decoration-thickness

Use the `decoration-[<value>]` syntax to set the text decoration thickness based on a completely custom value.

```html
<p class="decoration-[0.25rem] ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Resize element content to cover container with object-cover (HTML)

Source: https://tailwindcss.com/docs/object-fit

Use the `object-cover` utility to resize an element's content to cover its container.

```html
<img class="h-48 w-96 object-cover ..." src="/img/mountains.jpg" />
```

--------------------------------

### Whitespace Normal Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-normal` utility to cause text to wrap normally. Newlines and spaces will be collapsed.

```html
<p class="whitespace-normal">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Using Multiple Arguments in --value()

Source: https://tailwindcss.com/docs/adding-custom-styles

Pass multiple arguments to `--value()` to resolve them sequentially from left to right.

```css
@theme {
  --tab-size-github: 8;
}
@utility tab-* {
  tab-size: --value(--tab-size-*, integer, [integer]);
}
@utility opacity-* {
  opacity: calc(--value(integer) * 1%);
  opacity: --value(--opacity-*, [percentage]);
}
```

--------------------------------

### Using slashed zeroes

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `slashed-zero` utility to force a zero with a slash in fonts that support them.

```html
<p class="slashed-zero ...">0</p>
```

--------------------------------

### Enable Multiple OpenType Features

Source: https://tailwindcss.com/docs/font-feature-settings

Enable multiple OpenType features by separating them with commas within the `font-features` utility.

```html
<p class="font-features-['smcp','onum'] ...">This text uses small caps and oldstyle numbers.</p>
```

--------------------------------

### Define a Custom Variant with Shorthand Syntax

Source: https://tailwindcss.com/docs/adding-custom-styles

A concise way to define a custom variant when no complex nesting is required. This shorthand applies the provided selector directly.

```css
@custom-variant theme-midnight (&:where([data-theme="midnight"] *));
```

--------------------------------

### Container Query Ranges

Source: https://tailwindcss.com/docs/responsive-design

Target a specific range for container queries by stacking a regular container query variant with a max-width container query variant.

```html
<div class="@container">  <div class="flex flex-row @sm:@max-md:flex-col">    <!-- ... -->  </div></div>
```

--------------------------------

### Set Border-inline-start Color with Black

Source: https://tailwindcss.com/docs/border-color

Use `border-s-black` to set the inline-start border color to black using the predefined `--color-black` variable.

```css
border-inline-start-color: var(--color-black); /* #000 */
```

--------------------------------

### Preserve v3 Default Placeholder Color

Source: https://tailwindcss.com/docs/upgrade-guide

Add this CSS to your project to maintain the v3 default placeholder color.

```css
@layer base {
  input::placeholder,
  textarea::placeholder {
    color: var(--color-gray-400);
  }
}
```

--------------------------------

### Using a custom shrink value

Source: https://tailwindcss.com/docs/flex-shrink

Set the flex shrink factor using a custom value with the `shrink-[<value>]` syntax. This allows for precise control over shrinking behavior.

```html
<div class="shrink-[calc(100vw-var(--sidebar))] ...">  <!-- ... --></div>
```

--------------------------------

### Traditional CSS vs. Tailwind Variants

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Compares traditional CSS where a single class handles multiple states to Tailwind's approach using variants for conditional styling.

```html
<button class="btn">Save changes</button><style>
  .btn {
    background-color: var(--color-sky-500);
    &:hover {
      background-color: var(--color-sky-700);
    }
  }
</style>
```

--------------------------------

### Responsive Word Break Design

Source: https://tailwindcss.com/docs/word-break

Prefix a `word-break` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This allows for adaptive text wrapping based on viewport size.

```html
<p class="break-normal md:break-all ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Responsive Outline Color

Source: https://tailwindcss.com/docs/outline-color

Prefix an `outline-color` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This enables responsive design for outline colors.

```html
<div class="outline md:outline-blue-400 ...">  <!-- ... --></div>
```

--------------------------------

### Apply Basic Font Weight Classes in HTML

Source: https://tailwindcss.com/docs/font-weight

Use predefined Tailwind CSS classes like `font-light` or `font-bold` to control the font weight of an HTML element.

```html
<p class="font-light ...">The quick brown fox ...</p><p class="font-normal ...">The quick brown fox ...</p><p class="font-medium ...">The quick brown fox ...</p><p class="font-semibold ...">The quick brown fox ...</p><p class="font-bold ...">The quick brown fox ...</p>
```

--------------------------------

### Using Custom Line Height Values

Source: https://tailwindcss.com/docs/line-height

Set line height using arbitrary values with `leading-[<value>]` or CSS variables with `leading-(<custom-property>)`. The latter is a shorthand for `leading-[var(<custom-property>)]`.

```html
<p class="leading-[1.5] ...">  Lorem ipsum dolor sit amet...</p>
```

```html
<p class="leading-(--my-line-height) ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Import CSS Files with @import

Source: https://tailwindcss.com/docs/functions-and-directives

Use the @import directive to inline import CSS files, including Tailwind itself.

```css
@import "tailwindcss";
```

--------------------------------

### Custom grid column values with col-[<value>]

Source: https://tailwindcss.com/docs/grid-column

Set grid column sizes and locations using arbitrary values with utilities like `col-[<value>]`, `col-span-[<value>]`, `col-start-[<value>]`, and `col-end-[<value>]`.

```html
<div class="col-[16_/_span_16] ...">
  <!-- ... -->
</div>
```

--------------------------------

### Responsive underline offset with breakpoint variant

Source: https://tailwindcss.com/docs/text-underline-offset

Prefix underline-offset utilities with breakpoint variants like md: to apply the utility only at specific screen sizes and above.

```html
<p class="underline md:underline-offset-4 ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Custom Outline Width Value

Source: https://tailwindcss.com/docs/outline-width

Use the `outline-[<value>]` syntax for arbitrary outline widths, such as viewport units (`2vw`).

```html
<div class="outline-[2vw] ...">  <!-- ... --></div>
```

--------------------------------

### Stacking arbitrary variants

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Arbitrary variants can be stacked with built-in variants or with each other, allowing for complex conditional styling.

```html
<ul role="list">
  {#each items as item}
    <li class="[&.is-dragging]:active:cursor-grabbing">{item}</li>
  {/each}
</ul>
```

--------------------------------

### Applying Responsive Overscroll Behavior with Tailwind CSS

Source: https://tailwindcss.com/docs/overscroll-behavior

Prefix an `overscroll-behavior` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above.

```html
<div class="overscroll-auto md:overscroll-contain ...">  <!-- ... --></div>
```

--------------------------------

### Apply Mask Utilities

Source: https://tailwindcss.com/docs/mask-image

Utilize generated mask utility classes in your HTML to apply custom mask effects. Ensure the corresponding theme variables are defined.

```html
<div class="mask-radial-from-regal-blue">  <!-- ... --></div>
```

--------------------------------

### Set minimum block size with spacing scale

Source: https://tailwindcss.com/docs/min-block-size

Use `min-block-<number>` utilities to set a fixed minimum block size based on the spacing scale. Applied to elements within a container that also has a `block-20` class.

```html
<div class="block-20 ...">  <div class="min-block-80 ...">min-block-80</div>  <div class="min-block-64 ...">min-block-64</div>  <div class="min-block-48 ...">min-block-48</div>  <div class="min-block-40 ...">min-block-40</div>  <div class="min-block-32 ...">min-block-32</div></div>
```

--------------------------------

### Update Outline Utility Usage

Source: https://tailwindcss.com/docs/upgrade-guide

Replace `outline-none` with `outline-hidden` to reflect the change in how the utility handles outline styles.

```html
<input class="outline outline-2" /><input class="outline-2" />
```

```html
<input class="focus:outline-none" /><input class="focus:outline-hidden" />
```

--------------------------------

### Custom Skew Value

Source: https://tailwindcss.com/docs/skew

Use the `skew-[<value>]` syntax for arbitrary skew values, such as `skew-[3.142rad]`. This allows for precise control beyond predefined utilities.

```html
<img class="skew-[3.142rad] ..." src="/img/mountains.jpg" />
```

--------------------------------

### CSS for scrollbar-thumb-transparent utility

Source: https://tailwindcss.com/docs/scrollbar-color

Applies CSS to make the scrollbar thumb transparent.

```css
--tw-scrollbar-thumb: transparent; scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### Set caret color to Mist 50

Source: https://tailwindcss.com/docs/caret-color

Use the `caret-mist-50` class to set the caret color to a light mist shade. This utility allows for fine-grained control over the text cursor's appearance.

```html
<input class="caret-mist-50" />
```

--------------------------------

### Apply Styles for Hover, Focus, and Active States

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Utilize `hover:`, `focus:`, and `active:` variants to style elements differently when they are hovered over, focused on, or actively being clicked.

```html
<button class="bg-violet-500 hover:bg-violet-600 focus:outline-2 focus:outline-offset-2 focus:outline-violet-500 active:bg-violet-700 ...">
  Save changes
</button>
```

--------------------------------

### Balanced Text Wrapping

Source: https://tailwindcss.com/docs/text-wrap

Use the `text-balance` utility to distribute text evenly across each line. This is best suited for headings and may be limited by browsers to blocks of approximately 6 lines or less for performance reasons.

```html
<article>  <h3 class="text-balance">Beloved Manhattan soup stand closes</h3>  <p>New Yorkers are facing the winter chill...</p></article>
```

--------------------------------

### Restore v3 Button Cursor Behavior

Source: https://tailwindcss.com/docs/upgrade-guide

Add these base styles to CSS to restore the v3 default `cursor: pointer` behavior for buttons.

```css
@layer base {
  button:not(:disabled),
  [role="button"]:not(:disabled) {
    cursor: pointer;
  }
}
```

--------------------------------

### Apply horizontal padding with px-<number>

Source: https://tailwindcss.com/docs/padding

Use `px-<number>` utilities to control the horizontal padding of an element.

```html
<div class="px-8 ...">px-8</div>
```

--------------------------------

### Accessibility for Unstyled Lists

Source: https://tailwindcss.com/docs/preflight

Add a `role="list"` to unstyled lists to ensure they are announced correctly by screen readers.

```html
<ul role="list">
  <li>One</li>
  <li>Two</li>
  <li>Three</li>
</ul>
```

--------------------------------

### Using arbitrary accent color values

Source: https://tailwindcss.com/docs/accent-color

Apply an arbitrary accent color value directly using square bracket notation. This is useful for one-off color assignments or when the color is not predefined.

```html
<input class="accent-[<value>]" type="checkbox">
```

--------------------------------

### Importing Tailwind CSS Layers

Source: https://tailwindcss.com/docs/theme

This snippet shows the core imports when `tailwindcss` is included, defining the layers for theme, base, components, and utilities.

```css
@layer theme, base, components, utilities;@import "./theme.css" layer(theme);@import "./preflight.css" layer(base);@import "./utilities.css" layer(utilities);
```

--------------------------------

### Apply custom color utility class in HTML

Source: https://tailwindcss.com/docs/theme

After defining a custom color theme variable, use the corresponding utility class in your HTML to apply the color.

```html
<div class="bg-mint-500">  <!-- ... --></div>
```

--------------------------------

### Applying Same Styles for Multiple Variants

Source: https://tailwindcss.com/docs/adding-custom-styles

Apply the same styles for multiple variants by separating them with a comma within the @variant directive.

```css
.my-element {
  background: white;
  @variant hover, focus {
    background: black;
  }
}
```

--------------------------------

### Set font size with a custom value

Source: https://tailwindcss.com/docs/font-size

Use the `text-[<value>]` syntax to apply a custom font size not defined in the theme. This allows for precise control over text sizing.

```html
<p class="text-[14px] ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Applying Multiple Tailwind CSS Color Utilities

Source: https://tailwindcss.com/docs/colors

Shows how to apply various Tailwind CSS color utilities like `bg-*`, `border-*`, and `text-*` to style an element, including dark mode considerations.

```html
<div class="flex items-center gap-4 rounded-lg bg-white p-6 shadow-md outline outline-black/5 dark:bg-gray-800">  <span class="inline-flex shrink-0 rounded-full border border-pink-300 bg-pink-100 p-2 dark:border-pink-300/10 dark:bg-pink-400/10">    <svg class="size-6 stroke-pink-700 dark:stroke-pink-500"><!-- ... --></svg>  </span>  <div>    <p class="text-gray-700 dark:text-gray-400">      <span class="font-medium text-gray-950 dark:text-white">Tom Watson</span> mentioned you in      <span class="font-medium text-gray-950 dark:text-white">Logo redesign</span>    </p>    <time class="mt-1 block text-gray-500" datetime="9:37">9:37am</time>  </div></div>
```

--------------------------------

### Migrate Vite Configuration to Tailwind CSS v4 Plugin

Source: https://tailwindcss.com/docs/upgrade-guide

For Vite users, switch from the PostCSS plugin to the dedicated `@tailwindcss/vite` plugin for better performance and developer experience.

```typescript
import { defineConfig } from "vite"
import tailwindcss from "@tailwindcss/vite"
export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
});
```

--------------------------------

### Allowing Text to Wrap

Source: https://tailwindcss.com/docs/text-wrap

Use the `text-wrap` utility to allow overflowing text to wrap onto multiple lines at logical points.

```html
<article class="text-wrap">  <h3>Beloved Manhattan soup stand closes</h3>  <p>New Yorkers are facing the winter chill...</p></article>
```

--------------------------------

### Align items by baseline

Source: https://tailwindcss.com/docs/align-items

Use the `items-baseline` utility to align flex items so their baselines match.

```html
<div class="flex items-baseline ...">
  <div class="pt-2 pb-6">01</div>
  <div class="pt-8 pb-12">02</div>
  <div class="pt-12 pb-4">03</div>
</div>
```

--------------------------------

### Setting Custom Background Position with Arbitrary Values in HTML

Source: https://tailwindcss.com/docs/background-position

Apply a custom background position using the bg-position-[<value>] syntax for precise control over the background image's placement.

```html
<div class="bg-position-[center_top_1rem] ...">  <!-- ... --></div>
```

--------------------------------

### Responsive float design

Source: https://tailwindcss.com/docs/float

Control float behavior across different screen sizes by prefixing float utilities with responsive breakpoints, such as `md:float-left`.

```html
<img class="float-right md:float-left" src="/img/mountains.jpg" />
```

--------------------------------

### Define Custom Font Size with Line Height, Letter Spacing, and Font Weight

Source: https://tailwindcss.com/docs/font-size

You can provide default values for `line-height`, `letter-spacing`, and `font-weight` alongside the font size definition. This allows for more comprehensive customization of text appearance.

```css
@theme {  --text-tiny: 0.625rem;  --text-tiny--line-height: 1.5rem;   --text-tiny--letter-spacing: 0.125rem;   --text-tiny--font-weight: 500; }
```

--------------------------------

### justify-around

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-around` utility to distribute items evenly along the container's main axis with equal space on each side of each item.

```html
<div class="flex justify-around ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Generate Spacing Values with --spacing()

Source: https://tailwindcss.com/docs/functions-and-directives

Use the --spacing() function to generate a spacing value based on your theme's spacing scale. It accepts a multiplier.

```css
.my-element {
  margin: --spacing(4);
}
```

```css
.my-element {
  margin: calc(var(--spacing) * 4);
}
```

--------------------------------

### Implement Subgrid

Source: https://tailwindcss.com/docs/grid-template-columns

Use the grid-cols-subgrid utility to adopt column tracks defined by the parent grid. This is useful for aligning items within a complex grid structure.

```html
<div class="grid grid-cols-4 gap-4">  <div>01</div>  <!-- ... -->  <div>05</div>  <div class="col-span-3 grid grid-cols-subgrid gap-4">    <div class="col-start-2">06</div>  </div></div>
```

--------------------------------

### Custom Scale with CSS Variables

Source: https://tailwindcss.com/docs/scale

Use the `scale-(<custom-property>)` syntax to set the scale using CSS variables. This is a shorthand for `scale-[var(<custom-property>)]`.

```html
<img class="scale-(--my-scale) ..." src="/img/mountains.jpg" />
```

--------------------------------

### Skipping Snap Position Stops

Source: https://tailwindcss.com/docs/scroll-snap-stop

Use the `snap-normal` utility to enable a snap container to skip past possible scroll snap positions. This provides a more fluid scrolling experience where users are not forced to stop at every snap point.

```html
<div class="snap-x ...">
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-01.jpg" />
  </div>
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-02.jpg" />
  </div>
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-03.jpg" />
  </div>
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-04.jpg" />
  </div>
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-05.jpg" />
  </div>
  <div class="snap-center snap-normal ...">
    <img src="/img/vacation-06.jpg" />
  </div>
</div>
```

--------------------------------

### Allowing flex items to shrink

Source: https://tailwindcss.com/docs/flex-shrink

Use the `shrink` utility to allow a flex item to shrink if needed. This is the default behavior for flex items.

```html
<div class="flex ..."> <div class="h-14 w-14 flex-none ...">01</div> <div class="h-14 w-64 shrink ...">02</div> <div class="h-14 w-14 flex-none ...">03</div></div>
```

--------------------------------

### Conditional Styling for Pointer Coarseness

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use `pointer-coarse` and `pointer-fine` variants to apply styles based on the accuracy of the user's pointing device. Useful for optimizing UI for touchscreens.

```html
<fieldset aria-label="Choose a memory option">  <div class="flex items-center justify-between">    <div>RAM</div>    <a href="#"> See performance specs </a>  </div>  <div class="mt-4 grid grid-cols-6 gap-2 pointer-coarse:mt-6 pointer-coarse:grid-cols-3 pointer-coarse:gap-4">    <label class="p-2 pointer-coarse:p-4 ...">      <input type="radio" name="memory-option" value="4 GB" className="sr-only" />      <span>4 GB</span>    </label>    <!-- ... -->  </div></fieldset>
```

--------------------------------

### Setting flex grow with a CSS variable using Tailwind CSS `grow-(<custom-property>)`

Source: https://tailwindcss.com/docs/flex-grow

Apply a CSS variable as the flex grow factor using the `grow-(--my-grow)` syntax, which automatically wraps the variable in `var()`.

```html
<div class="grow-(--my-grow) ...">  <!-- ... --></div>
```

--------------------------------

### Center and Pad Tailwind CSS Container

Source: https://tailwindcss.com/docs/max-width

Add `mx-auto` and `px-<number>` utilities to a `container` to center it and apply horizontal padding.

```html
<div class="container mx-auto px-4">  <!-- ... --></div>
```

--------------------------------

### Configure Astro to use Tailwind CSS Vite plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/astro

Add the `@tailwindcss/vite` plugin to your Astro configuration file to enable Tailwind CSS processing.

```javascript
// @ts-checkimport { defineConfig } from "astro/config";import tailwindcss from "@tailwindcss/vite";// https://astro.build/configexport default defineConfig({
  vite: {
    plugins: [tailwindcss()],
  },
});
```

--------------------------------

### Using Custom @supports Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply custom `@supports` variants in your project for cleaner and more concise conditional styling based on CSS feature support.

```html
<div class="supports-grid:grid">  <!-- ... --></div>
```

--------------------------------

### Using Custom Gradient Value

Source: https://tailwindcss.com/docs/background-image

Set gradient properties using custom values with utilities like `bg-linear-[<value>]`. Supports CSS variables with `bg-linear-(<custom-property>)` syntax.

```html
<div class="bg-linear-[25deg,red_5%,yellow_60%,lime_90%,teal] ...">  <!-- ... --></div>
```

```html
<div class="bg-linear-(--my-gradient) ...">  <!-- ... --></div>
```

--------------------------------

### Set Fixed Maximum Width with Tailwind CSS

Source: https://tailwindcss.com/docs/max-width

Use `max-w-<number>` utilities to apply a fixed maximum width based on the spacing scale.

```html
<div class="w-full max-w-96 ...">max-w-96</div><div class="w-full max-w-80 ...">max-w-80</div><div class="w-full max-w-64 ...">max-w-64</div><div class="w-full max-w-48 ...">max-w-48</div><div class="w-full max-w-40 ...">max-w-40</div><div class="w-full max-w-32 ...">max-w-32</div><div class="w-full max-w-24 ...">max-w-24</div>
```

--------------------------------

### Use Custom Grid Column Value

Source: https://tailwindcss.com/docs/grid-template-columns

Set grid columns using a custom value with the grid-cols-[<value>] syntax. This allows for precise control over column sizing.

```html
<div class="grid-cols-[200px_minmax(900px,_1fr)_100px] ...">  <!-- ... --></div>
```

--------------------------------

### Using a Custom Font Utility Class in HTML

Source: https://tailwindcss.com/docs/theme

This HTML snippet shows how to apply the custom `font-script` utility class, defined in the extended theme, to an HTML element.

```html
<p class="font-script">This will use the Great Vibes font family.</p>
```

--------------------------------

### Using custom accent color properties

Source: https://tailwindcss.com/docs/accent-color

Set a custom accent color by referencing a CSS custom property using the accent-color utility. This allows for dynamic color theming.

```html
<input class="accent-(<custom-property>)" type="checkbox">
```

--------------------------------

### Create arbitrary group variants with selectors

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Define one-off `group-*` variants dynamically by providing a custom selector in square brackets, like `group-[.selector]`. Use the `&` character to control the placement of `.group` within the generated selector for complex targeting.

```html
<div class="group is-published">
  <div class="hidden group-[.is-published]:block">
    Published
  </div>
</div>
```

```html
<div class="group">
  <div class="group-[:nth-of-type(3)_&]:block">
    <!-- ... -->
  </div>
</div>
```

--------------------------------

### Target a Single Breakpoint Range

Source: https://tailwindcss.com/docs/responsive-design

To target a single breakpoint, stack a responsive variant with the `max-*` variant for the next breakpoint. This ensures the style applies only at the specified breakpoint.

```html
<div class="md:max-lg:flex">  <!-- ... --></div>
```

--------------------------------

### Basic Zoom Utilities

Source: https://tailwindcss.com/docs/zoom

Use `zoom-<number>` utilities like `zoom-75` and `zoom-125` to scale an element using the CSS `zoom` property. These utilities apply predefined zoom levels.

```html
<div class="zoom-75 ...">  <!-- ... --></div><div class="zoom-100 ...">  <!-- ... --></div><div class="zoom-125 ...">  <!-- ... --></div>
```

--------------------------------

### Set box-sizing to content-box

Source: https://tailwindcss.com/docs/box-sizing

Use the `box-content` utility to add borders and padding on top of the element's specified width or height. This makes the element larger than its defined dimensions.

```html
<div class="box-content size-32 border-4 p-4 ...">  <!-- ... --></div>
```

--------------------------------

### Responsive Translate with Tailwind CSS Breakpoints

Source: https://tailwindcss.com/docs/translate

Prefix a `translate` utility with a breakpoint variant like `md:` to apply the translation only at specific screen sizes and above.

```html
<img class="translate-45 md:translate-60 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Mark Utilities as Important

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, use `!` at the end of a class name to mark it as important. The old prefix position is deprecated.

```html
<div class="flex! bg-red-500! hover:bg-red-600/50!">  <!-- ... --></div>
```

--------------------------------

### Integrate compiled Tailwind CSS into HTML

Source: https://tailwindcss.com/docs/installation/tailwind-cli

Link the generated 'output.css' file in your HTML document's '<head>' section to apply Tailwind's utility classes for styling.

```HTML
<!doctype html><html><head>  <meta charset="UTF-8">  <meta name="viewport" content="width=device-width, initial-scale=1.0">  <link href="./output.css" rel="stylesheet"></head><body>  <h1 class="text-3xl font-bold underline">    Hello world!  </h1></body></html>
```

--------------------------------

### Avoid dynamic class names in HTML

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Tailwind cannot understand string concatenation or interpolation. Ensure all class names are complete and statically present in the source files to be detected.

```html
<div class="text-{{ error ? 'red' : 'green' }}-600"></div>
```

--------------------------------

### Use a customized animation utility in Tailwind CSS

Source: https://tailwindcss.com/docs/animation

Apply a custom animation utility, such as `animate-wiggle`, after defining it in the theme configuration.

```html
<div class="animate-wiggle">  <!-- ... --></div>
```

--------------------------------

### Configure Theme Variable Generation

Source: https://tailwindcss.com/docs/preflight

When importing Tailwind CSS files individually, use `theme(static)` or `theme(inline)` on the `theme.css` import to control how generated theme variables are handled.

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/theme.css" layer(theme) theme(static);
@import "tailwindcss/utilities.css" layer(utilities);
```

--------------------------------

### Style Dialog Backdrop

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `backdrop` variant to style the background of a native `<dialog>` element. This allows customization of the overlay when a dialog is open.

```html
<dialog class="backdrop:bg-gray-50">  <form method="dialog">    <!-- ... -->  </form></dialog>
```

--------------------------------

### Scale down element content to fit container with object-scale-down (HTML)

Source: https://tailwindcss.com/docs/object-fit

Use the `object-scale-down` utility to display an element's content at its original size but scale it down to fit its container if necessary.

```html
<img class="h-48 w-96 object-scale-down ..." src="/img/mountains.jpg" />
```

--------------------------------

### Brightness with CSS custom property

Source: https://tailwindcss.com/docs/filter-brightness

Use the brightness-(<custom-property>) syntax as shorthand for brightness-[var(<custom-property>)], which automatically wraps the custom property in the var() function.

```html
<img class="brightness-(--my-brightness) ..." src="/img/mountains.jpg" />
```

--------------------------------

### Stretch content in Grid with Tailwind CSS

Source: https://tailwindcss.com/docs/place-content

Use `place-content-stretch` to make grid items stretch to fill their grid areas along both the inline and block axes.

```html
<div class="grid h-48 grid-cols-2 place-content-stretch gap-4 ...">  <div>01</div>  <div>02</div>  <div>03</div>  <div>04</div></div>
```

--------------------------------

### Set border color to lime shades

Source: https://tailwindcss.com/docs/border-color

Apply various shades of lime to the border using classes like `border-lime-50` through `border-lime-300`.

```html
<div class="border border-lime-50 ..."></div>
<div class="border border-lime-100 ..."></div>
<div class="border border-lime-200 ..."></div>
<div class="border border-lime-300 ..."></div>
```

--------------------------------

### Enable PostCSS Loader in webpack.config.js

Source: https://tailwindcss.com/docs/installation/framework-guides/symfony

Configure Webpack Encore to use the PostCSS Loader for processing CSS.

```javascript
Encore
  .enablePostCssLoader();
```

--------------------------------

### Set custom outline offset using arbitrary values

Source: https://tailwindcss.com/docs/outline-offset

Apply a specific, custom outline offset using the `outline-offset-[<value>]` arbitrary value syntax.

```html
<div class="outline-offset-[2vw] ...">  <!-- ... --></div>
```

--------------------------------

### Configure Tailwind CSS Vite plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/laravel/vite

Add the `@tailwindcss/vite` plugin to your Vite configuration file.

```typescript
import { defineConfig } from 'vite'import tailwindcss from '@tailwindcss/vite'export default defineConfig({  plugins: [    tailwindcss(),    // …  ],})
```

--------------------------------

### Use Tailwind utility classes in HTML

Source: https://tailwindcss.com/docs/installation/framework-guides/parcel

Link the compiled CSS file in your HTML document and apply Tailwind utility classes to style content.

```html
<!doctype html><html>  <head>    <meta charset="UTF-8" />    <meta name="viewport" content="width=device-width, initial-scale=1.0" />    <link href="./index.css" type="text/css" rel="stylesheet" />  </head>  <body>    <h1 class="text-3xl font-bold underline">      Hello world!    </h1>  </body></html>
```

--------------------------------

### End-start logical border radius with arbitrary value

Source: https://tailwindcss.com/docs/border-radius

Apply end-start border-radius using an arbitrary CSS value.

```css
border-end-start-radius: <value>;
```

--------------------------------

### Custom text-decoration-thickness with CSS variable

Source: https://tailwindcss.com/docs/text-decoration-thickness

For CSS variables, use the `decoration-(length:<custom-property>)` syntax. This is a shorthand for `decoration-[length:var(<custom-property>)]`.

```html
<p class="decoration-(length:--my-decoration-thickness) ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Custom z-index value

Source: https://tailwindcss.com/docs/z-index

Use the z-[<value>] syntax for custom z-index values.

```html
<div class="z-[calc(var(--index)+1)] ...">  <!-- ... --></div>
```

--------------------------------

### Background Size: CSS Variable

Source: https://tailwindcss.com/docs/background-size

Use `bg-size-(<custom-property>)` as a shorthand for `bg-size-[var(<custom-property>)]` to apply background sizes defined by CSS variables.

```html
<div class="bg-size-(--my-image-size) ...">  <!-- ... --></div>
```

--------------------------------

### Apply custom scrollbar colors with arbitrary values

Source: https://tailwindcss.com/docs/scrollbar-color

Use bracket notation to set scrollbar colors with completely custom hex or color values.

```html
<div class="scrollbar-thumb-[#0f766e] scrollbar-track-[#ccfbf1] ...">  <!-- ... --></div>
```

--------------------------------

### Responsive text overflow variants

Source: https://tailwindcss.com/docs/text-overflow

Apply text overflow utilities responsively by prefixing them with breakpoint variants like `md:`. This allows the overflow behavior to change at different screen sizes.

```html
<p class="text-ellipsis md:text-clip ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Customize caret color theme

Source: https://tailwindcss.com/docs/caret-color

Use the `--color-*` theme variables to customize the color utilities in your project. The `caret-regal-blue` utility can then be used in your markup.

```css
@theme {  --color-regal-blue: #243c5a; }
```

```html
<textarea class="caret-regal-blue"></textarea>
```

--------------------------------

### Set box-sizing to border-box

Source: https://tailwindcss.com/docs/box-sizing

Use the `box-border` utility to include borders and padding when defining an element's height or width. This is the default in Tailwind's preflight styles.

```html
<div class="box-border size-32 border-4 p-4 ...">  <!-- ... --></div>
```

--------------------------------

### Setting Inline Border Color to Yellow Shades

Source: https://tailwindcss.com/docs/border-color

Illustrates setting the inline border color using different shades of yellow, ranging from `yellow-50` to `yellow-950`.

```html
<!-- Set border-inline-color to yellow-50 -->
<div class="border-x-yellow-50"></div>

<!-- Set border-inline-color to yellow-100 -->
<div class="border-x-yellow-100"></div>

<!-- Set border-inline-color to yellow-200 -->
<div class="border-x-yellow-200"></div>

<!-- Set border-inline-color to yellow-300 -->
<div class="border-x-yellow-300"></div>

<!-- Set border-inline-color to yellow-400 -->
<div class="border-x-yellow-400"></div>

<!-- Set border-inline-color to yellow-500 -->
<div class="border-x-yellow-500"></div>

<!-- Set border-inline-color to yellow-600 -->
<div class="border-x-yellow-600"></div>

<!-- Set border-inline-color to yellow-700 -->
<div class="border-x-yellow-700"></div>

<!-- Set border-inline-color to yellow-800 -->
<div class="border-x-yellow-800"></div>

<!-- Set border-inline-color to yellow-900 -->
<div class="border-x-yellow-900"></div>

<!-- Set border-inline-color to yellow-950 -->
<div class="border-x-yellow-950"></div>
```

--------------------------------

### Reduced motion support with motion-reduce variant

Source: https://tailwindcss.com/docs/transition-duration

Use the motion-reduce variant to set duration-0 for users who prefer reduced motion, ensuring accessibility compliance.

```html
<button type="button" class="duration-300 motion-reduce:duration-0 ...">  <!-- ... --></button>
```

--------------------------------

### Update PostCSS Configuration for Tailwind CSS v4

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, the PostCSS plugin is in a separate package. Remove `postcss-import` and `autoprefixer` if no longer needed, as v4 handles imports and vendor prefixing automatically.

```javascript
export default {
  plugins: {
    "postcss-import": {},
    tailwindcss: {},
    autoprefixer: {},
    "@tailwindcss/postcss": {},
  },
};
```

--------------------------------

### Arbitrary Grid Columns

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Specify complex grid layouts using arbitrary values within square brackets.

```html
<div class="grid grid-cols-[24rem_2.5rem_minmax(0,1fr)]">  <!-- ... --></div>
```

--------------------------------

### Apply Subpixel Antialiasing with Tailwind CSS

Source: https://tailwindcss.com/docs/font-smoothing

Use the `subpixel-antialiased` utility class to apply subpixel font smoothing. This can result in sharper text rendering on displays with a known pixel layout.

```html
<p class="subpixel-antialiased ...">The quick brown fox ...</p>
```

--------------------------------

### place-self-auto alignment in grid

Source: https://tailwindcss.com/docs/place-self

Align a grid item based on the container's place-items property using the place-self-auto class.

```html
<div class="grid grid-cols-3 gap-4 ...">
  <div>01</div>
  <div class="place-self-auto ...">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### Supporting Negative Values

Source: https://tailwindcss.com/docs/adding-custom-styles

Register separate positive and negative utilities into distinct declarations to support negative values.

```css
@utility inset-* {
  inset: --spacing(--value(integer));
  inset: --value([percentage], [length]);
}
@utility -inset-* {
  inset: --spacing(--value(integer) * -1);
  inset: calc(--value([percentage], [length]) * -1);
}
```

--------------------------------

### Set Fixed Block Size with Spacing Scale

Source: https://tailwindcss.com/docs/block-size

Use `block-<number>` utilities like `block-24` and `block-64` to set an element to a fixed block size based on the spacing scale.

```html
<div class="block-96 ...">block-96</div><div class="block-80 ...">block-80</div><div class="block-64 ...">block-64</div><div class="block-48 ...">block-48</div><div class="block-40 ...">block-40</div><div class="block-32 ...">block-32</div><div class="block-24 ...">block-24</div>
```

--------------------------------

### Arbitrary Properties with CSS Variables

Source: https://tailwindcss.com/docs/adding-custom-styles

Use arbitrary properties to manage CSS variables, allowing them to change under different conditions like responsive variants.

```html
<div class="[--scroll-offset:56px] lg:[--scroll-offset:44px]">  <!-- ... --></div>
```

--------------------------------

### Logical Scroll Margin Properties (Block)

Source: https://tailwindcss.com/docs/scroll-margin

Use `scroll-mbs-6` and `scroll-mbe-6` for `scroll-margin-block-start` and `scroll-margin-block-end`. These map to top or bottom based on writing mode.

```html
<div class="snap-y ...">  <div class="snap-start scroll-mbs-6 ...">    <!-- ... -->  </div></div>
```

--------------------------------

### Responsive Scroll Margin

Source: https://tailwindcss.com/docs/scroll-margin

Prefix scroll margin utilities with a breakpoint variant like `md:` to apply them only at specific screen sizes and above. This enables adaptive layouts.

```html
<div class="scroll-m-8 md:scroll-m-0 ...">  <!-- ... --></div>
```

--------------------------------

### Apply responsive outline offset with Tailwind CSS

Source: https://tailwindcss.com/docs/outline-offset

Conditionally apply an outline offset based on screen size by prefixing the utility with a breakpoint variant like `md:`.

```html
<div class="outline md:outline-offset-2 ...">  <!-- ... --></div>
```

--------------------------------

### Matching Viewport Inline Size

Source: https://tailwindcss.com/docs/inline-size

Use the `inline-screen` utility to make an element span the entire inline size of the viewport.

```html
<div class="inline-screen">  <!-- ... --></div>
```

--------------------------------

### Define a custom color theme variable

Source: https://tailwindcss.com/docs/theme

Use the `@theme` directive in your CSS to define a new color variable, making it available for utility classes like `bg-mint-500`.

```css
@import "tailwindcss";@theme {  --color-mint-500: oklch(0.72 0.11 178);}
```

--------------------------------

### Basic transform-origin utilities

Source: https://tailwindcss.com/docs/transform-origin

Use utilities like `origin-center`, `origin-top-left`, and `origin-bottom` to set an element's transform origin. These utilities directly map to CSS `transform-origin` values.

```html
<img class="origin-center rotate-45 ..." src="/img/mountains.jpg" /><img class="origin-top-left rotate-12 ..." src="/img/mountains.jpg" /><img class="origin-bottom -rotate-12 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive transform-origin

Source: https://tailwindcss.com/docs/transform-origin

Prefix a `transform-origin` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This allows for responsive adjustments to the transform origin.

```html
<img class="origin-center md:origin-top ..." src="/img/mountains.jpg" />
```

--------------------------------

### Snapping to the end

Source: https://tailwindcss.com/docs/scroll-snap-align

Use the `snap-end` utility to snap an element to its end when being scrolled inside a snap container.

```html
<div class="snap-x ...">  <div class="snap-end ...">    <img src="/img/vacation-01.jpg" />  </div>  <div class="snap-end ...">    <img src="/img/vacation-02.jpg" />  </div>  <div class="snap-end ...">    <img src="/img/vacation-03.jpg" />  </div>  <div class="snap-end ...">    <img src="/img/vacation-04.jpg" />  </div>  <div class="snap-end ...">    <img src="/img/vacation-05.jpg" />  </div>  <div class="snap-end ...">    <img src="/img/vacation-06.jpg" />  </div></div>
```

--------------------------------

### Handle Gradient Variants in Dark Mode

Source: https://tailwindcss.com/docs/upgrade-guide

Use `via-none` to explicitly unset three-stop gradients when overriding with variants in specific states like dark mode.

```html
<div class="bg-gradient-to-r from-red-500 to-yellow-400 dark:from-blue-500">
  <!-- ... --></div>
```

```html
<div class="bg-linear-to-r from-red-500 via-orange-400 to-yellow-400 dark:via-none dark:from-blue-500 dark:to-teal-400">
  <!-- ... --></div>
```

--------------------------------

### Custom Backdrop Invert Value

Source: https://tailwindcss.com/docs/backdrop-filter-invert

Use the `backdrop-invert-[<value>]` syntax to set the backdrop inversion based on a completely custom value. This allows for precise control over the inversion percentage.

```html
<div class="backdrop-invert-[.25] ...">  <!-- ... --></div>
```

--------------------------------

### Set Border Color with Mauve Palette

Source: https://tailwindcss.com/docs/border-color

Apply border colors from the mauve color palette using these utilities. Each class sets the `border-inline-color` with a specific oklch value for the mauve scale.

```html
`border-x-mauve-50`| `border-inline-color: var(--color-mauve-50); /* oklch(98.5% 0 0) */`
```

```html
`border-x-mauve-100`| `border-inline-color: var(--color-mauve-100); /* oklch(96% 0.003 325.6) */`
```

```html
`border-x-mauve-200`| `border-inline-color: var(--color-mauve-200); /* oklch(92.2% 0.005 325.62) */`
```

```html
`border-x-mauve-300`| `border-inline-color: var(--color-mauve-300); /* oklch(86.5% 0.012 325.68) */`
```

```html
`border-x-mauve-400`| `border-inline-color: var(--color-mauve-400); /* oklch(71.1% 0.019 323.02) */`
```

```html
`border-x-mauve-500`| `border-inline-color: var(--color-mauve-500); /* oklch(54.2% 0.034 322.5) */`
```

```html
`border-x-mauve-600`| `border-inline-color: var(--color-mauve-600); /* oklch(43.5% 0.029 321.78) */`
```

```html
`border-x-mauve-700`| `border-inline-color: var(--color-mauve-700); /* oklch(36.4% 0.029 323.89) */`
```

```html
`border-x-mauve-800`| `border-inline-color: var(--color-mauve-800); /* oklch(26.3% 0.024 320.12) */`
```

```html
`border-x-mauve-900`| `border-inline-color: var(--color-mauve-900); /* oklch(21.2% 0.019 322.12) */`
```

```html
`border-x-mauve-950`| `border-inline-color: var(--color-mauve-950); /* oklch(14.5% 0.008 326) */`
```

--------------------------------

### Removing Box Shadows

Source: https://tailwindcss.com/docs/box-shadow

Remove existing box shadows, inset shadows, or rings using the `shadow-none`, `inset-shadow-none`, `ring-0`, and `inset-ring-0` utilities.

```html
<div class="shadow-none ..."></div>
```

--------------------------------

### Style File Input Buttons

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

The `file` variant allows styling the button within file input elements. Customize appearance like margins, borders, background, padding, and text styles.

```html
<input  type="file"  class="file:mr-4 file:rounded-full file:border-0 file:bg-violet-50 file:px-4 file:py-2 file:text-sm file:font-semibold file:text-violet-700 hover:file:bg-violet-100 dark:file:bg-violet-600 dark:file:text-violet-100 dark:hover:file:bg-violet-500 ..."/>
```

--------------------------------

### Stretch items

Source: https://tailwindcss.com/docs/align-items

Use the `items-stretch` utility to make flex items fill the container's cross axis.

```html
<div class="flex items-stretch ...">
  <div class="py-4">01</div>
  <div class="py-12">02</div>
  <div class="py-8">03</div>
</div>
```

--------------------------------

### Arbitrary variants with at-rules

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Arbitrary variants can incorporate at-rules like `@media` or `@supports` for conditional styling based on media queries or feature support.

```html
<div class="flex [@supports(display:grid)]:grid">
  <!-- ... -->
</div>
```

--------------------------------

### Apply Custom Width with Tailwind CSS `w-[<value>]`

Source: https://tailwindcss.com/docs/width

Apply a custom width value to an element using the arbitrary value syntax `w-[<value>]`.

```html
<div class="w-[5px] ...">  <!-- ... --></div>
```

--------------------------------

### Safelisting a specific utility

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Use `@source inline()` to force Tailwind to generate a specific class name that might not be present in your content files.

```css
@import "tailwindcss";@source inline("underline");
```

--------------------------------

### Custom Invert Value

Source: https://tailwindcss.com/docs/filter-invert

Use the `invert-[<value>]` syntax to apply a custom inversion percentage, such as `.25`.

```html
<img class="invert-[.25] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Tailwind CSS Custom Backdrop Saturate with CSS Variables

Source: https://tailwindcss.com/docs/backdrop-filter-saturate

For CSS variables, use the `backdrop-saturate-(<custom-property>)` syntax. This is a shorthand for `backdrop-saturate-[var(<custom-property>)]` that automatically adds the `var()` function.

```html
<div class="backdrop-saturate-(--my-backdrop-saturation) ...">  <!-- ... --></div>
```

--------------------------------

### Inset Box Shadows

Source: https://tailwindcss.com/docs/box-shadow

Apply an inset box shadow using utilities like `inset-shadow-xs` and `inset-shadow-sm`. The opacity of inset shadows can be adjusted with an opacity modifier, e.g., `inset-shadow-sm/50`.

```html
<div class="inset-shadow-2xs ..."></div><div class="inset-shadow-xs ..."></div><div class="inset-shadow-sm ..."></div>
```

--------------------------------

### Set fixed maximum height with max-h-<number>

Source: https://tailwindcss.com/docs/max-height

Apply fixed maximum heights using spacing scale utilities like max-h-24 and max-h-64.

```html
<div class="h-96 ...">  <div class="h-full max-h-80 ...">max-h-80</div>  <div class="h-full max-h-64 ...">max-h-64</div>  <div class="h-full max-h-48 ...">max-h-48</div>  <div class="h-full max-h-40 ...">max-h-40</div>  <div class="h-full max-h-32 ...">max-h-32</div>  <div class="h-full max-h-24 ...">max-h-24</div></div>
```

--------------------------------

### Define Custom Font Size Variable

Source: https://tailwindcss.com/docs/font-size

Use the `--text-*` theme variables to customize font size utilities in your project. This sets the base font size for the `text-tiny` utility.

```css
@theme {  --text-tiny: 0.625rem; }
```

--------------------------------

### Styling Visited Links

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles to links that have already been visited using the `visited` variant.

```html
<a href="https://seinfeldquotes.com" class="text-blue-600 visited:text-purple-600 ..."> Inspiration </a>
```

--------------------------------

### Define Custom Design Tokens with @theme

Source: https://tailwindcss.com/docs/functions-and-directives

Use the @theme directive to define project-specific design tokens like fonts, colors, and breakpoints. Learn more in the theme variables documentation.

```css
@theme {  --font-display: "Satoshi", "sans-serif";  --breakpoint-3xl: 120rem;  --color-avocado-100: oklch(0.99 0 0);  --color-avocado-200: oklch(0.98 0.04 113.22);  --color-avocado-300: oklch(0.94 0.11 115.03);  --color-avocado-400: oklch(0.92 0.19 114.08);  --color-avocado-500: oklch(0.84 0.18 117.33);  --color-avocado-600: oklch(0.53 0.12 118.34);  --ease-fluid: cubic-bezier(0.3, 0, 0, 1);  --ease-snappy: cubic-bezier(0.2, 0, 0, 1);  /* ... */}
```

--------------------------------

### justify-stretch

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-stretch` utility to allow auto-sized content items to fill the available space along the container's main axis.

```html
<div class="grid grid-cols-[4rem_auto_4rem] justify-stretch ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Customizing font-feature-settings and font-variation-settings

Source: https://tailwindcss.com/docs/font-family

Set default font-feature-settings and font-variation-settings for a custom font family in the theme.

```css
@theme {  --font-display: "Oswald", sans-serif;  --font-display--font-feature-settings: "cv02", "cv03", "cv04", "cv11";   --font-display--font-variation-settings: "opsz" 32; }
```

--------------------------------

### Responsive border radius

Source: https://tailwindcss.com/docs/border-radius

Prefix border-radius utilities with breakpoint variants like md: to apply them conditionally at specific screen sizes.

```html
<div class="rounded md:rounded-lg ...">  <!-- ... --></div>
```

--------------------------------

### Backdrop brightness with CSS custom property

Source: https://tailwindcss.com/docs/backdrop-filter-brightness

Use the backdrop-brightness-(<custom-property>) syntax as shorthand for backdrop-brightness-[var(<custom-property>)], which automatically wraps the custom property in the var() function.

```html
<div class="backdrop-brightness-(--my-backdrop-brightness) ...">  <!-- ... --></div>
```

--------------------------------

### Enable Horizontal Resizing Only

Source: https://tailwindcss.com/docs/resize

Use the `resize-x` utility to permit an element to be resized exclusively along the horizontal axis. This allows for adjustments left and right, but not up and down.

```html
<textarea class="resize-x rounded-md ..."></textarea>
```

--------------------------------

### Use Custom Mask Value

Source: https://tailwindcss.com/docs/mask-image

Set a mask image using a completely custom value with utilities like `mask-linear-[<value>]` or `mask-radial-[<value>]`. This allows for precise control over the mask's appearance.

```html
<div class="mask-linear-[70deg,transparent_10%,black,transparent_80%] ...">  <!-- ... --></div>
```

--------------------------------

### Set Custom Block Size Value

Source: https://tailwindcss.com/docs/block-size

Use the `block-[<value>]` syntax to set the block size based on a completely custom value.

```html
<div class="block-[32rem] ...">  <!-- ... --></div>
```

--------------------------------

### Import global styles as reference in Vue component styles

Source: https://tailwindcss.com/docs/compatibility

Use @reference in scoped style blocks to make theme variables available for @apply. This pattern applies to Vue, Svelte, and Astro component files.

```Vue
<template>
  <button><slot /></button>
</template>

<style scoped>
  @reference "../app.css";
  button {
    @apply bg-blue-500;
  }
</style>
```

--------------------------------

### Using lining figures

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `lining-nums` utility to use numeric glyphs that are aligned by their baseline in fonts that support them.

```html
<p class="lining-nums ...">1234567890</p>
```

--------------------------------

### Applying Custom Percentage Font Stretch in HTML

Source: https://tailwindcss.com/docs/font-stretch

Use the `font-stretch-[<value>]` syntax to set a custom font width using an arbitrary percentage value.

```html
<p class="font-stretch-[66.66%] ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Set font size using CSS variables

Source: https://tailwindcss.com/docs/font-size

Apply font sizes defined by CSS variables using the `text-(length:--my-text-size)` syntax. This is a shorthand for `text-[length:var(--my-text-size)]`.

```html
<p class="text-(length:--my-text-size) ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Basic Border Colors

Source: https://tailwindcss.com/docs/border-color

Use utilities like `border-rose-500` and `border-lime-100` to control the border color of an element.

```html
<div class="border-4 border-indigo-500 ..."></div><div class="border-4 border-purple-500 ..."></div><div class="border-4 border-sky-500 ..."></div>
```

--------------------------------

### Setting border-left color with specific values

Source: https://tailwindcss.com/docs/border-color

Control the border-left color using keywords like 'inherit', 'current', 'transparent', or specific colors like 'black' and 'white'.

```html
<div class="border-l-inherit"></div>
<div class="border-l-current"></div>
<div class="border-l-transparent"></div>
<div class="border-l-black"></div>
<div class="border-l-white"></div>
```

--------------------------------

### place-self-end alignment in grid

Source: https://tailwindcss.com/docs/place-self

Align a grid item to the end on both axes using the place-self-end class.

```html
<div class="grid grid-cols-3 gap-4 ...">
  <div>01</div>
  <div class="place-self-end ...">02</div>
  <div>03</div>
  <div>04</div>
  <div>05</div>
  <div>06</div>
</div>
```

--------------------------------

### Custom list marker with CSS variable

Source: https://tailwindcss.com/docs/list-style-type

Use the list-(<custom-property>) syntax as shorthand for list-[var(<custom-property>)] to reference a CSS variable.

```html
<ol class="list-(--my-marker) ...">
  <!-- ... -->
</ol>
```

--------------------------------

### Scale Utility on Hover

Source: https://tailwindcss.com/docs/scale

Apply scale utilities conditionally in a hover state using the `hover:` variant prefix. This allows for interactive scaling effects.

```html
<img class="scale-95 hover:scale-120 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Background Color with Custom Property

Source: https://tailwindcss.com/docs/background-color

Applies a background color using a custom CSS property. Replace `<custom-property>` with your CSS variable name.

```html
<div class="bg-(<custom-property>)">...</div>
```

--------------------------------

### Generated CSS for Breakpoint Variant

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Shows the generated CSS for a responsive utility, using a media query to apply styles only when the screen width meets the specified breakpoint.

```css
.sm\:grid-cols-3 {
  @media (width >= 40rem) {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
}
```

--------------------------------

### CSS Variable Backdrop Contrast

Source: https://tailwindcss.com/docs/backdrop-filter-contrast

For CSS variables, use the `backdrop-contrast-(<custom-property>)` syntax. This is a shorthand for `backdrop-contrast-[var(<custom-property>)]`.

```html
<div class="backdrop-contrast-(--my-backdrop-contrast) ...">  <!-- ... --></div>
```

--------------------------------

### Snapping to the center

Source: https://tailwindcss.com/docs/scroll-snap-align

Use the `snap-center` utility to snap an element to its center when being scrolled inside a snap container.

```html
<div class="snap-x ...">  <div class="snap-center ...">    <img src="/img/vacation-01.jpg" />  </div>  <div class="snap-center ...">    <img src="/img/vacation-02.jpg" />  </div>  <div class="snap-center ...">    <img src="/img/vacation-03.jpg" />  </div>  <div class="snap-center ...">    <img src="/img/vacation-04.jpg" />  </div>  <div class="snap-center ...">    <img src="/img/vacation-05.jpg" />  </div>  <div class="snap-center ...">    <img src="/img/vacation-06.jpg" />  </div></div>
```

--------------------------------

### Apply color-scheme with dark mode variant

Source: https://tailwindcss.com/docs/color-scheme

Use the dark: variant to apply different color-scheme utilities based on dark mode state. Prefix the utility with dark:* to conditionally apply the scheme only when dark mode is active.

```html
<html class="scheme-light dark:scheme-dark ...">  <!-- ... --></html>
```

--------------------------------

### Apply Custom Hue Rotation with Arbitrary Values in Tailwind CSS

Source: https://tailwindcss.com/docs/filter-hue-rotate

Set a custom hue rotation using arbitrary values with the `hue-rotate-[<value>]` syntax.

```html
<img class="hue-rotate-[3.142rad] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Responsive Scrollbar Gutter Design

Source: https://tailwindcss.com/docs/scrollbar-gutter

Prefix a `scrollbar-gutter` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. Learn more about using variants in the variants documentation.

```html
<div class="scrollbar-gutter-auto md:scrollbar-gutter-stable ...">  <!-- ... --></div>
```

--------------------------------

### Set border color to white

Source: https://tailwindcss.com/docs/border-color

Use `border-white` to set the border color to white.

```html
<div class="border border-white ..."></div>
```

--------------------------------

### Stretch element content to fit container with object-fill (HTML)

Source: https://tailwindcss.com/docs/object-fit

Use the `object-fill` utility to stretch an element's content to fit its container.

```html
<img class="h-48 w-96 object-fill ..." src="/img/mountains.jpg" />
```

--------------------------------

### Apply Outline Styles with Tailwind CSS

Source: https://tailwindcss.com/docs/outline-style

Use `outline-solid`, `outline-dashed`, `outline-dotted`, and `outline-double` utilities to set various outline styles on elements.

```html
<button class="outline-2 outline-offset-2 outline-solid ...">Button A</button><button class="outline-2 outline-offset-2 outline-dashed ...">Button B</button><button class="outline-2 outline-offset-2 outline-dotted ...">Button C</button><button class="outline-3 outline-offset-2 outline-double ...">Button D</button>
```

--------------------------------

### Complex Grid Layout with Inline Styles

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Apply a complex grid layout using inline styles for intricate arbitrary values that are hard to represent as class names.

```html
<div class="grid-[2fr_max(0,var(--gutter-width))_calc(var(--gutter-width)+10px)]"><div style="grid-template-columns: 2fr max(0, var(--gutter-width)) calc(var(--gutter-width) + 10px)">
  <!-- ... -->
</div>
```

--------------------------------

### Backdrop Invert with CSS Variables

Source: https://tailwindcss.com/docs/backdrop-filter-invert

For CSS variables, use the `backdrop-invert-(<custom-property>)` syntax. This is a shorthand for `backdrop-invert-[var(<custom-property>)]`.

```html
<div class="backdrop-invert-(--my-backdrop-inversion) ...">  <!-- ... --></div>
```

--------------------------------

### Customizing Container Sizes

Source: https://tailwindcss.com/docs/responsive-design

Define custom container sizes using the `--container-*` theme variables in your CSS. This allows for specific breakpoints not included by default.

```css
@import "tailwindcss";@theme {  --container-8xl: 96rem;}
```

--------------------------------

### Combining Arbitrary Values with Modifiers

Source: https://tailwindcss.com/docs/adding-custom-styles

Combine arbitrary values with interactive (hover) and responsive (lg) modifiers for dynamic styling.

```html
<div class="top-[117px] lg:top-[344px]">  <!-- ... --></div>
```

--------------------------------

### Basic Letter Spacing in HTML

Source: https://tailwindcss.com/docs/letter-spacing

Apply different letter spacing utilities like `tracking-tight`, `tracking-normal`, and `tracking-wide` to paragraph elements.

```html
<p class="tracking-tight ...">The quick brown fox ...</p><p class="tracking-normal ...">The quick brown fox ...</p><p class="tracking-wide ...">The quick brown fox ...</p>
```

--------------------------------

### Using Custom Letter Spacing Utility

Source: https://tailwindcss.com/docs/letter-spacing

Apply a newly defined custom letter spacing utility, such as `tracking-tightest`, to an HTML element after it has been added to the theme.

```html
<p class="tracking-tightest">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Using the default scrollbar width

Source: https://tailwindcss.com/docs/scrollbar-width

Use the `scrollbar-auto` utility to use the browser's default scrollbar width.

```html
<div class="scrollbar-auto overflow-auto ...">  <!-- ... --></div>
```

--------------------------------

### Set Bottom Border Color with Stone Palette

Source: https://tailwindcss.com/docs/border-color

Apply `border-b-{color}-{shade}` utilities for the stone color palette to define bottom border colors.

```html
border-b-stone-50
border-b-stone-100
border-b-stone-200
border-b-stone-300
border-b-stone-400
border-b-stone-500
border-b-stone-600
border-b-stone-700
border-b-stone-800
border-b-stone-900
border-b-stone-950
```

--------------------------------

### Configure Vite plugin for Tailwind CSS

Source: https://tailwindcss.com/docs/installation/framework-guides/adonisjs

Add the `@tailwindcss/vite` plugin to your `vite.config.ts` file to enable Tailwind CSS processing.

```TypeScript
import { defineConfig } from 'vite'import adonisjs from '@adonisjs/vite/client'import tailwindcss from '@tailwindcss/vite'export default defineConfig({
  plugins: [
    tailwindcss(),
    adonisjs({
      // …
    }),
  ],
})
```

--------------------------------

### Apply custom breakpoint variant in HTML

Source: https://tailwindcss.com/docs/theme

Use the variant generated from a custom breakpoint theme variable to apply utilities conditionally based on screen size.

```html
<div class="3xl:grid-cols-6 grid grid-cols-2 md:grid-cols-4">  <!-- ... --></div>
```

--------------------------------

### justify-evenly

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-evenly` utility to distribute items evenly along the container's main axis with equal space around each item, accounting for doubled space between items.

```html
<div class="flex justify-evenly ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Apply Custom Shadow Color to Drop Shadow

Source: https://tailwindcss.com/docs/filter-drop-shadow

Once a custom color variable is defined, you can use it with the `drop-shadow-*` utility class to apply a specific color to your shadows.

```html
<svg class="drop-shadow-regal-blue">  <!-- ... --></svg>
```

--------------------------------

### CSS Variable for Contrast

Source: https://tailwindcss.com/docs/filter-contrast

Utilize CSS variables for contrast with the `contrast-(<custom-property>)` syntax, which is a shorthand for `contrast-[var(<custom-property>)]`.

```html
<img class="contrast-(--my-contrast) ..." src="/img/mountains.jpg" />
```

--------------------------------

### Using Custom CSS Properties for Text Color

Source: https://tailwindcss.com/docs/text-color

Apply text colors using custom CSS properties. This allows for dynamic color changes or integration with CSS variables.

```html
text-(<custom-property>)
```

--------------------------------

### Configure Tailwind CSS Vite plugin in Nuxt

Source: https://tailwindcss.com/docs/installation/framework-guides/nuxt

Integrate the `@tailwindcss/vite` plugin into your `nuxt.config.ts` file under the `vite.plugins` array.

```typescript
import tailwindcss from "@tailwindcss/vite";export default defineNuxtConfig({  compatibilityDate: "2025-07-15",  devtools: { enabled: true },  vite: {    plugins: [      tailwindcss(),    ],  },});
```

--------------------------------

### Styling on Focus

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles when an element receives focus using the `focus` variant.

```html
<input class="border-gray-300 focus:border-blue-400 ..." />
```

--------------------------------

### Customize Box Shadow - Tailwind CSS

Source: https://tailwindcss.com/docs/box-shadow

Use the --shadow-* theme variables to customize the box shadow utilities. This allows for custom shadow definitions that can be applied via utility classes.

```css
@theme {  --shadow-3xl: 0 35px 35px rgba(0, 0, 0, 0.25); }
```

```html
<div class="shadow-3xl">  <!-- ... --></div>
```

--------------------------------

### Responsive Font Smoothing with Tailwind CSS

Source: https://tailwindcss.com/docs/font-smoothing

Apply font smoothing utilities responsively by prefixing them with a breakpoint variant like `md:`. This allows you to change font smoothing behavior at different screen sizes.

```html
<p class="antialiased md:subpixel-antialiased ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Apply Custom Numeric Font Weight in HTML

Source: https://tailwindcss.com/docs/font-weight

Set a specific numeric font weight using arbitrary values with the `font-[<value>]` syntax.

```html
<p class="font-[1000] ...">  Lorem ipsum dolor sit amet...</p>
```

--------------------------------

### Custom Drop Shadow Value

Source: https://tailwindcss.com/docs/filter-drop-shadow

Apply a drop shadow with a completely custom value using the `drop-shadow-[<value>]` syntax.

```html
<svg class="drop-shadow-[0_35px_35px_rgba(0,0,0,0.25)] ...">  <!-- ... --></svg>
```

--------------------------------

### Generated CSS for Dark Mode Variant

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Illustrates the generated CSS for the 'dark:' variant, utilizing a media query for 'prefers-color-scheme: dark' to apply styles conditionally.

```css
.dark\:bg-gray-800 {
  @media (prefers-color-scheme: dark) {
    background-color: var(--color-gray-800);
  }
}
```

--------------------------------

### Update Border and Divide Utilities to Specify Color

Source: https://tailwindcss.com/docs/upgrade-guide

In v4, border and divide utilities default to `currentColor`. Explicitly set a color like `border-gray-200` to maintain v3 behavior.

```html
<div class="border border-gray-200 px-2 py-3 ...">  <!-- ... --></div>
```

--------------------------------

### Using oldstyle figures

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `oldstyle-nums` utility to use numeric glyphs where some numbers have descenders in fonts that support them.

```html
<p class="oldstyle-nums ...">1234567890</p>
```

--------------------------------

### Break All Words

Source: https://tailwindcss.com/docs/word-break

Use the `break-all` utility to add line breaks whenever necessary, without trying to preserve whole words. This can be useful for preventing overflow with long, unbroken strings.

```html
<p class="break-all">The longest word in any of the major English language dictionaries is pneumonoultramicroscopicsilicovolcanoconiosis, a word that refers to a lung disease contracted from the inhalation of very fine silica particles, specifically from a volcano; medically, it is the same as silicosis.</p>
```

--------------------------------

### Configure Tailwind CSS Vite Plugin

Source: https://tailwindcss.com/docs/installation/framework-guides/solidjs

Add the @tailwindcss/vite plugin to your Vite configuration file.

```typescript
import { defineConfig } from 'vite';import solidPlugin from 'vite-plugin-solid';import tailwindcss from '@tailwindcss/vite';export default defineConfig({  plugins: [    tailwindcss(),    solidPlugin(),  ],  server: {    port: 3000,  },  build: {    target: 'esnext',  },});
```

--------------------------------

### Apply basic outline offset with Tailwind CSS

Source: https://tailwindcss.com/docs/outline-offset

Use predefined `outline-offset-<number>` utilities to control the spacing between an element's outline and its border.

```html
<button class="outline-2 outline-offset-0 ...">Button A</button><button class="outline-2 outline-offset-2 ...">Button B</button><button class="outline-2 outline-offset-4 ...">Button C</button>
```

--------------------------------

### Basic Grayscale Filters

Source: https://tailwindcss.com/docs/filter-grayscale

Use utilities like `grayscale` and `grayscale-75` to control the amount of grayscale effect applied to an element. `grayscale-0` removes the effect entirely.

```html
<img class="grayscale-0 ..." src="/img/mountains.jpg" /><img class="grayscale-25 ..." src="/img/mountains.jpg" /><img class="grayscale-50 ..." src="/img/mountains.jpg" /><img class="grayscale ..." src="/img/mountains.jpg" />
```

--------------------------------

### Custom Box Shadow with CSS Variable

Source: https://tailwindcss.com/docs/box-shadow

Apply a box shadow using a custom CSS property. This allows for dynamic shadow values defined elsewhere.

```html
<div class="shadow-(<custom-property>)"></div>
```

--------------------------------

### Style direct children with the * variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `*` variant to apply styles to direct children of an element. This is useful when you need to style children you don't have direct control over.

```html
<div>
  <h2>Categories<h2>
  <ul class="*:rounded-full *:border *:border-sky-100 *:bg-sky-50 *:px-2 *:py-0.5 dark:text-sky-300 dark:*:border-sky-500/15 dark:*:bg-sky-500/10 ...">
    <li>Sales</li>
    <li>Marketing</li>
    <li>SEO</li>
    <!-- ... -->
  </ul>
</div>
```

--------------------------------

### Order flex items first or last with Tailwind CSS

Source: https://tailwindcss.com/docs/order

Apply `order-first` or `order-last` utilities to position specific flex or grid items at the beginning or end of the container, respectively.

```html
<div class="flex justify-between ...">  <div class="order-last ...">01</div>  <div class="...">02</div>  <div class="order-first ...">03</div></div>
```

--------------------------------

### Match Viewport Height with Tailwind CSS

Source: https://tailwindcss.com/docs/height

Use the `h-screen` utility to make an element span the entire height of the viewport.

```html
<div class="h-screen">  <!-- ... --></div>
```

--------------------------------

### Style First Line and First Letter

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Style the first line of text with the `first-line` variant and the first letter with the `first-letter` variant. These are useful for stylistic emphasis in typography.

```html
<div class="text-gray-700">  <p    class="first-letter:float-left first-letter:mr-3 first-letter:text-7xl first-letter:font-bold first-letter:text-gray-900 first-line:tracking-widest first-line:uppercase"  >    Well, let me tell you something, funny boy. Y'know that little stamp, the one that says "New York Public Library"?  </p>  <p class="mt-6">Well that may not mean anything to you, but that means a lot to me. One whole hell of a lot.</p></div>
```

--------------------------------

### Applying arbitrary border colors

Source: https://tailwindcss.com/docs/border-color

Use this utility to apply any arbitrary border color value. This is useful for one-off color assignments or when a specific color is not available in the predefined scales.

```css
& > :not(:last-child) {   border-color: <value>; }
```

--------------------------------

### Set inline-end border color to lime

Source: https://tailwindcss.com/docs/border-color

Use these utilities to set the inline-end border color to different shades of lime. This demonstrates the application of lighter, brighter colors to borders.

```html
<div class="border-e-lime-50">...</div>
<div class="border-e-lime-100">...</div>
<div class="border-e-lime-200">...</div>
<div class="border-e-lime-300">...</div>
<div class="border-e-lime-400">...</div>
```

--------------------------------

### Arbitrary CSS Variables

Source: https://tailwindcss.com/docs/styling-with-utility-classes

Define arbitrary CSS properties and values, useful for setting CSS variables.

```html
<div class="[--gutter-width:1rem] lg:[--gutter-width:2rem]">  <!-- ... --></div>
```

--------------------------------

### Add Linear Gradient

Source: https://tailwindcss.com/docs/background-image

Combine directional utilities like `bg-linear-to-r` with color stop utilities to create linear gradients. Specify angles for custom gradient directions.

```html
<div class="h-14 bg-linear-to-r from-cyan-500 to-blue-500"></div><div class="h-14 bg-linear-to-t from-sky-500 to-indigo-500"></div><div class="h-14 bg-linear-to-bl from-violet-500 to-fuchsia-500"></div><div class="h-14 bg-linear-65 from-purple-500 to-pink-500"></div>
```

--------------------------------

### Apply vertical padding with py-<number>

Source: https://tailwindcss.com/docs/padding

Use `py-<number>` utilities to control the vertical padding of an element.

```html
<div class="py-8 ...">py-8</div>
```

--------------------------------

### Custom Scale Value

Source: https://tailwindcss.com/docs/scale

Apply a completely custom scale value using the `scale-[<value>]` syntax. This allows for precise control over element scaling.

```html
<img class="scale-[1.7] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Add Radial Gradient

Source: https://tailwindcss.com/docs/background-image

Use `bg-radial` with optional position modifiers like `bg-radial-[at_50%_75%]` and color stop utilities to create radial gradients. The `from-` and `to-` utilities define the gradient colors.

```html
<div class="size-18 rounded-full bg-radial from-pink-400 from-40% to-fuchsia-700"></div><div class="size-18 rounded-full bg-radial-[at_50%_75%] from-sky-200 via-blue-400 to-indigo-900 to-90%"></div><div class="size-18 rounded-full bg-radial-[at_25%_25%] from-white to-zinc-900 to-75%"></div>
```

--------------------------------

### Styling First and Last List Items

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `first` and `last` variants to remove top and bottom padding from the first and last list items respectively.

```html
<ul role="list">  {#each people as person}    <!-- Remove top/bottom padding when first/last child -->    <li class="flex py-4 first:pt-0 last:pb-0">      <img class="h-10 w-10 rounded-full" src={person.imageUrl} alt="" />      <div class="ml-3 overflow-hidden">        <p class="text-sm font-medium text-gray-900 dark:text-white">{person.name}</p>        <p class="truncate text-sm text-gray-500 dark:text-gray-400">{person.email}</p>      </div>    </li>  {/each}</ul>
```

--------------------------------

### Logical properties for text direction

Source: https://tailwindcss.com/docs/border-radius

Use logical property utilities like rounded-s-lg and rounded-se-xl to apply border radius based on text direction (LTR or RTL).

```html
<div dir="ltr">  <div class="rounded-s-lg ..."></div></div><div dir="rtl">  <div class="rounded-s-lg ..."></div></div>
```

--------------------------------

### Custom Backdrop Opacity Value

Source: https://tailwindcss.com/docs/backdrop-filter-opacity

Apply custom backdrop filter opacity values using the `backdrop-opacity-[<value>]` syntax. This allows for precise control beyond the predefined utility classes.

```html
<div class="backdrop-opacity-[.15] ...">  <!-- ... --></div>
```

--------------------------------

### Apply Custom Gap Values in Tailwind CSS

Source: https://tailwindcss.com/docs/gap

Use `gap-[<value>]` syntax to define a custom gap size, such as `gap-[10vw]`, for precise spacing control.

```html
<div class="gap-[10vw] ...">  <!-- ... --></div>
```

--------------------------------

### Applying Custom will-change Value

Source: https://tailwindcss.com/docs/will-change

Use the `will-change-[<value>]` syntax to set the `will-change` property with a custom value, such as `top,left`. This allows for more specific animation optimizations.

```html
<div class="will-change-[top,left] ...">  <!-- ... --></div>
```

--------------------------------

### Differentiating peers with unique names

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

When using multiple peers, assign unique names using `peer/{name}` classes to style based on a specific peer's state. Variants like `peer-checked/{name}` target the element based on the named peer's checked state.

```html
<fieldset>
  <legend>Published status</legend>
  <input id="draft" class="peer/draft" type="radio" name="status" checked />
  <label for="draft" class="peer-checked/draft:text-sky-500">Draft</label>
  <input id="published" class="peer/published" type="radio" name="status" />
  <label for="published" class="peer-checked/published:text-sky-500">Published</label>
  <div class="hidden peer-checked/draft:block">Drafts are only visible to administrators.</div>
  <div class="hidden peer-checked/published:block">Your post will be publicly visible on your site.</div>
</fieldset>
```

--------------------------------

### CSS for scrollbar-track-sky-400 utility

Source: https://tailwindcss.com/docs/scrollbar-color

Applies the `scrollbar-track-sky-400` utility to set the scrollbar track color to sky-400.

```css
--tw-scrollbar-track: var(--color-sky-400); /* oklch(74.6% 0.16 232.661) */ scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### Proximity Scroll Snapping

Source: https://tailwindcss.com/docs/scroll-snap-type

Use `snap-x` and `snap-proximity` to allow a snap container to rest on snap points close in proximity. Ensure children have snap alignment.

```html
<div class="snap-x snap-proximity ...">
  <div class="snap-center ...">
    <img src="/img/vacation-01.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-02.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-03.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-04.jpg" />
  </div>
  <div class="snap-center ...">
    <img src="/img/vacation-05.jpg" />
  </div>
</div>
```

--------------------------------

### Set custom order value with bracket syntax in Tailwind CSS

Source: https://tailwindcss.com/docs/order

Define a custom order value using arbitrary value syntax, `order-[<value>]`, for dynamic or complex ordering scenarios.

```html
<div class="order-[min(var(--total-items),10)] ...">  <!-- ... --></div>
```

--------------------------------

### Customize Drop Shadow Value

Source: https://tailwindcss.com/docs/filter-drop-shadow

Use the `--drop-shadow-*` theme variables to customize the drop shadow utilities in your project. This allows for fine-grained control over shadow appearance.

```css
@theme {  --drop-shadow-3xl: 0 35px 35px rgba(0, 0, 0, 0.25); }
```

--------------------------------

### Right Border Color Utilities (Mist)

Source: https://tailwindcss.com/docs/border-color

Apply mist colors to the right border using these utilities. Covers shades from 50 to 950.

```html
border-r-mist-50
```

```css
border-right-color: var(--color-mist-50); /* oklch(98.7% 0.002 197.1) */
```

```html
border-r-mist-100
```

```css
border-right-color: var(--color-mist-100); /* oklch(96.3% 0.002 197.1) */
```

```html
border-r-mist-200
```

```css
border-right-color: var(--color-mist-200); /* oklch(92.5% 0.005 214.3) */
```

```html
border-r-mist-300
```

```css
border-right-color: var(--color-mist-300); /* oklch(87.2% 0.007 219.6) */
```

```html
border-r-mist-400
```

```css
border-right-color: var(--color-mist-400); /* oklch(72.3% 0.014 214.4) */
```

```html
border-r-mist-500
```

```css
border-right-color: var(--color-mist-500); /* oklch(56% 0.021 213.5) */
```

```html
border-r-mist-600
```

```css
border-right-color: var(--color-mist-600); /* oklch(45% 0.017 213.2) */
```

```html
border-r-mist-700
```

```css
border-right-color: var(--color-mist-700); /* oklch(37.8% 0.015 216) */
```

```html
border-r-mist-800
```

```css
border-right-color: var(--color-mist-800); /* oklch(27.5% 0.011 216.9) */
```

```html
border-r-mist-900
```

```css
border-right-color: var(--color-mist-900); /* oklch(21.8% 0.008 223.9) */
```

```html
border-r-mist-950
```

```css
border-right-color: var(--color-mist-950); /* oklch(14.8% 0.004 228.8) */
```

--------------------------------

### Tailwind CSS Responsive Backdrop Saturate

Source: https://tailwindcss.com/docs/backdrop-filter-saturate

Prefix a `backdrop-filter: saturate()` utility with a breakpoint variant like `md:` to apply the utility only at medium screen sizes and above. This enables responsive adjustments to backdrop saturation.

```html
<div class="backdrop-saturate-50 md:backdrop-saturate-150 ...">  <!-- ... --></div>
```

--------------------------------

### Whitespace No Wrap Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-nowrap` utility to prevent text from wrapping. Newlines and spaces will be collapsed.

```html
<p class="overflow-auto whitespace-nowrap">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Set Bottom Border Color with Olive Palette

Source: https://tailwindcss.com/docs/border-color

Apply `border-b-{color}-{shade}` utilities for the olive color palette to control bottom border colors.

```html
border-b-olive-50
border-b-olive-100
border-b-olive-200
```

--------------------------------

### Configure Tailwind CSS Vite plugin

Source: https://tailwindcss.com/docs/installation/using-vite

Add the `@tailwindcss/vite` plugin to your Vite configuration file to enable Tailwind CSS processing.

```typescript
import { defineConfig } from 'vite'import tailwindcss from '@tailwindcss/vite'export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
})
```

--------------------------------

### Set Bottom Border Color with Zinc Palette

Source: https://tailwindcss.com/docs/border-color

Apply `border-b-{color}-{shade}` utilities for the zinc color palette to control bottom border colors.

```html
border-b-zinc-50
border-b-zinc-100
border-b-zinc-200
border-b-zinc-300
border-b-zinc-400
border-b-zinc-500
border-b-zinc-600
border-b-zinc-700
border-b-zinc-800
border-b-zinc-900
border-b-zinc-950
```

--------------------------------

### Custom Grayscale Value

Source: https://tailwindcss.com/docs/filter-grayscale

Use the `grayscale-[<value>]` syntax to set the grayscale based on a completely custom value, such as `0.5`.

```html
<img class="grayscale-[0.5] ..." src="/img/mountains.jpg" />
```

--------------------------------

### Use Tailwind CSS classes in an Astro component

Source: https://tailwindcss.com/docs/installation/framework-guides/astro

Import the global stylesheet into an Astro component and apply Tailwind's utility classes to HTML elements.

```astro
---import "../styles/global.css";---
<h1 class="text-3xl font-bold underline">
  Hello world!</h1>
```

--------------------------------

### Set Custom Transform Values

Source: https://tailwindcss.com/docs/transform

Use the `transform-[<value>]` syntax for arbitrary transform values like `matrix()`. For CSS variables, use the `transform-(<custom-property>)` shorthand.

```html
<div class="transform-[matrix(1,2,3,4,5,6)] ...">  <!-- ... --></div>
```

```html
<div class="transform-(--my-transform) ...">  <!-- ... --></div>
```

--------------------------------

### Custom Zoom Value

Source: https://tailwindcss.com/docs/zoom

Use the `zoom-[<value>]` syntax to set the zoom based on a completely custom value. This allows for precise control over the scaling factor.

```html
<div class="zoom-[1.1] ...">  <!-- ... --></div>
```

--------------------------------

### Use table-fixed for Fixed Column Widths

Source: https://tailwindcss.com/docs/table-layout

Use the `table-fixed` utility to enforce fixed widths for table columns, ignoring cell content. This ensures a consistent layout even with long text.

```html
<table class="table-fixed">  <thead>    <tr>      <th>Song</th>      <th>Artist</th>      <th>Year</th>    </tr>  </thead>  <tbody>    <tr>      <td>The Sliding Mr. Bones (Next Stop, Pottersville)</td>      <td>Malcolm Lockyer</td>      <td>1961</td>    </tr>    <tr>      <td>Witchy Woman</td>      <td>The Eagles</td>      <td>1972</td>    </tr>    <tr>      <td>Shining Star</td>      <td>Earth, Wind, and Fire</td>      <td>1975</td>    </tr>  </tbody></table>
```

--------------------------------

### Styling Target Element

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles to an element if its ID matches the current URL fragment using the `target` variant.

```html
<div id="about" class="target:shadow-lg ...">  <!-- ... --></div>
```

--------------------------------

### Using proportional figures

Source: https://tailwindcss.com/docs/font-variant-numeric

Use the `proportional-nums` utility to use numeric glyphs that have proportional widths in fonts that support them.

```html
<p class="proportional-nums ...">12121</p><p class="proportional-nums ...">90909</p>
```

--------------------------------

### Enable gatsby-plugin-postcss in gatsby-config.js

Source: https://tailwindcss.com/docs/installation/framework-guides/gatsby

Add `gatsby-plugin-postcss` to the plugins array in your Gatsby configuration file.

```JavaScript
module.exports = {  plugins: [    'gatsby-plugin-postcss',    // ...  ],}
```

--------------------------------

### Custom min-height value

Source: https://tailwindcss.com/docs/min-height

Set a minimum height using a completely custom pixel value with the min-h-[<value>] syntax.

```html
<div class="min-h-[220px] ...">  <!-- ... --></div>
```

--------------------------------

### Setting border-bottom color with custom properties

Source: https://tailwindcss.com/docs/border-color

Apply border-bottom color using CSS custom properties. This allows for dynamic color control through variables.

```html
<div class="border-b-[<custom-property>]"></div>
```

--------------------------------

### Set Custom Maximum Width with Arbitrary Values in Tailwind CSS

Source: https://tailwindcss.com/docs/max-width

Use the `max-w-[<value>]` syntax to define a maximum width with a completely custom value.

```html
<div class="max-w-[220px] ...">  <!-- ... --></div>
```

--------------------------------

### Applying text decoration on hover

Source: https://tailwindcss.com/docs/text-decoration-line

Prefix a `text-decoration-line` utility with a variant like `hover:*` to only apply the utility in that state.

```html
<p>The <a href="..." class="no-underline hover:underline ...">quick brown fox</a> jumps over the lazy dog.</p>
```

--------------------------------

### Responsive Align Content in Tailwind CSS

Source: https://tailwindcss.com/docs/align-content

Apply `align-content` utilities conditionally at different breakpoints using responsive prefixes like `md:` to adjust alignment based on screen size.

```html
<div class="grid content-start md:content-around ...">  <!-- ... --></div>
```

--------------------------------

### Translate elements using the spacing scale

Source: https://tailwindcss.com/docs/translate

Apply `translate-<number>` utilities to translate an element on both X and Y axes based on the configured spacing scale.

```html
<img class="-translate-6 ..." src="/img/mountains.jpg" /><img class="translate-2 ..." src="/img/mountains.jpg" /><img class="translate-8 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Arbitrary peer variants with '&' for selector positioning

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `&` character within arbitrary peer variant selectors to control the placement of the `.peer` class in the generated CSS. This offers more precise control over selector composition.

```html
<div>
  <input type="text" class="peer" />
  <div class="hidden peer-[:nth-of-type(3)_&]:block">
    <!-- ... -->
  </div>
</div>
```

--------------------------------

### Style Optional Inputs with :optional Variant

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Apply styles to input elements that are marked as optional.

```html
<input class="border optional:border-red-500 ..." />
```

--------------------------------

### Handle spaces in pseudo-element content

Source: https://tailwindcss.com/docs/content

Replace spaces with underscores when setting arbitrary string values for pseudo-element content. For literal underscores, escape them with a backslash.

```html
<p class="before:content-['Hello_World'] ..."></p>
```

```html
<p class="before:content-['Hello\_World']"></p>
```

--------------------------------

### End-end logical border radius with preset sizes

Source: https://tailwindcss.com/docs/border-radius

Apply border-radius to the end-end corner using logical properties. Supports sizes from xs (0.125rem) to 4xl (2rem).

```css
border-end-end-radius: var(--radius-xs); /* 0.125rem (2px) */
```

```css
border-end-end-radius: var(--radius-sm); /* 0.25rem (4px) */
```

```css
border-end-end-radius: var(--radius-md); /* 0.375rem (6px) */
```

```css
border-end-end-radius: var(--radius-lg); /* 0.5rem (8px) */
```

```css
border-end-end-radius: var(--radius-xl); /* 0.75rem (12px) */
```

```css
border-end-end-radius: var(--radius-2xl); /* 1rem (16px) */
```

```css
border-end-end-radius: var(--radius-3xl); /* 1.5rem (24px) */
```

```css
border-end-end-radius: var(--radius-4xl); /* 2rem (32px) */
```

--------------------------------

### Define a custom font theme variable

Source: https://tailwindcss.com/docs/theme

Add a new font family by defining a theme variable in the `--font-*` namespace to create a corresponding utility class.

```css
@import "tailwindcss";@theme {  --font-poppins: Poppins, sans-serif;}
```

--------------------------------

### Tailwind CSS Custom Backdrop Saturate Value

Source: https://tailwindcss.com/docs/backdrop-filter-saturate

Use the `backdrop-saturate-[<value>]` syntax to set the backdrop saturation based on a completely custom value. This allows for precise control beyond predefined options.

```html
<div class="backdrop-saturate-[.25] ...">  <!-- ... --></div>
```

--------------------------------

### justify-between

Source: https://tailwindcss.com/docs/justify-content

Use the `justify-between` utility to distribute items evenly along the container's main axis with equal space between each item.

```html
<div class="flex justify-between ...">  <div>01</div>  <div>02</div>  <div>03</div></div>
```

--------------------------------

### Whitespace Break Spaces Utility

Source: https://tailwindcss.com/docs/white-space

Use the `whitespace-break-spaces` utility to preserve newlines and spaces. White space at the end of lines will not hang, but will wrap to the next line.

```html
<p class="whitespace-break-spaces">Hey everyone!It's almost 2022       and we still don't know if there             are aliens living among us, or do we? Maybe the person writing this is an alien.You will never know.</p>
```

--------------------------------

### Responsive Rotation in Tailwind CSS

Source: https://tailwindcss.com/docs/rotate

Apply rotation utilities responsively by prefixing them with a breakpoint variant, such as `md:rotate-60`. This ensures the utility only applies at medium screen sizes and above.

```html
<img class="rotate-45 md:rotate-60 ..." src="/img/mountains.jpg" />
```

--------------------------------

### Enable Tailwind watcher in dev.exs

Source: https://tailwindcss.com/docs/installation/framework-guides/phoenix

Add Tailwind to the watchers list in development configuration to automatically rebuild CSS on file changes.

```elixir
watchers: [  # Start the esbuild watcher by calling Esbuild.install_and_run(:default, args)  esbuild: {Esbuild, :install_and_run, [:myproject, ~w(--sourcemap=inline --watch)]},  tailwind: {Tailwind, :install_and_run, [:myproject, ~w(--watch)]}]
```

--------------------------------

### Set border-y color to lime shades

Source: https://tailwindcss.com/docs/border-color

Apply various shades of lime to the top and bottom borders using the `border-y-{color}-{shade}` utility. This includes shades from 50 to 950.

```html
<div class="border-y-lime-50">...</div>
```

```html
<div class="border-y-lime-100">...</div>
```

```html
<div class="border-y-lime-200">...</div>
```

```html
<div class="border-y-lime-300">...</div>
```

```html
<div class="border-y-lime-400">...</div>
```

```html
<div class="border-y-lime-500">...</div>
```

```html
<div class="border-y-lime-600">...</div>
```

```html
<div class="border-y-lime-700">...</div>
```

```html
<div class="border-y-lime-800">...</div>
```

```html
<div class="border-y-lime-900">...</div>
```

```html
<div class="border-y-lime-950">...</div>
```

--------------------------------

### Disabling Automatic Detection and Explicitly Registering Sources

Source: https://tailwindcss.com/docs/detecting-classes-in-source-files

Completely disable automatic source detection with source(none) to manually register all desired source paths. Ensures each stylesheet only includes necessary classes.

```css
@import "tailwindcss" source(none);@source "../admin";@source "../shared";
```

--------------------------------

### Conditional Styling for Unsupported Features

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `not-supports-[...]` variant to apply styles when a specific CSS feature is not supported by the browser. This is the inverse of the `supports-[...]` variant.

```html
<div class="not-supports-[display:grid]:flex">  <!-- ... --></div>
```

--------------------------------

### CSS for scrollbar-track-sky-600 utility

Source: https://tailwindcss.com/docs/scrollbar-color

Applies the `scrollbar-track-sky-600` utility to set the scrollbar track color to sky-600.

```css
--tw-scrollbar-track: var(--color-sky-600); /* oklch(58.8% 0.158 241.966) */ scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### Add Custom Utilities with @utility

Source: https://tailwindcss.com/docs/functions-and-directives

Use the @utility directive to add custom utilities that work with variants like hover, focus, and lg. Learn more about registering custom utilities in the adding custom utilities documentation.

```css
@utility tab-4 {
  tab-size: 4;
}
```

--------------------------------

### Custom Box Shadow Value

Source: https://tailwindcss.com/docs/box-shadow

Apply a box shadow with a completely custom value using bracket notation, such as `shadow-[0_35px_35px_rgba(0,0,0,0.25)]`. This allows for precise control over shadow properties.

```html
<div class="shadow-[0_35px_35px_rgba(0,0,0,0.25)] ...">  <!-- ... --></div>
```

--------------------------------

### Left-side border radius with preset sizes

Source: https://tailwindcss.com/docs/border-radius

Apply border-radius to both top-left and bottom-left corners using predefined size variables. Use rounded-l-3xl for 1.5rem (24px) or rounded-l-4xl for 2rem (32px).

```css
border-top-left-radius: var(--radius-3xl); /* 1.5rem (24px) */ border-bottom-left-radius: var(--radius-3xl); /* 1.5rem (24px) */
```

```css
border-top-left-radius: var(--radius-4xl); /* 2rem (32px) */ border-bottom-left-radius: var(--radius-4xl); /* 2rem (32px) */
```

--------------------------------

### Set border-y color to emerald shades

Source: https://tailwindcss.com/docs/border-color

Apply various shades of emerald to the top and bottom borders using the `border-y-{color}-{shade}` utility. This includes shades from 50 to 200.

```html
<div class="border-y-emerald-50">...</div>
```

```html
<div class="border-y-emerald-100">...</div>
```

```html
<div class="border-y-emerald-200">...</div>
```

--------------------------------

### CSS variable min-inline-size utility

Source: https://tailwindcss.com/docs/min-inline-size

Use the `min-inline-(<custom-property>)` syntax as a shorthand for `min-inline-[var(<custom-property>)]` to automatically add the `var()` function.

```html
<div class="min-inline-(--my-min-inline-size) ...">  <!-- ... --></div>
```

--------------------------------

### Collapsing table borders with border-collapse

Source: https://tailwindcss.com/docs/border-collapse

Apply the border-collapse utility to a table element to combine adjacent cell borders into a single border. Include border utilities on the table and individual cells to define border styling.

```html
<table class="border-collapse border border-gray-400 ...">
  <thead>
    <tr>
      <th class="border border-gray-300 ...">State</th>
      <th class="border border-gray-300 ...">City</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class="border border-gray-300 ...">Indiana</td>
      <td class="border border-gray-300 ...">Indianapolis</td>
    </tr>
    <tr>
      <td class="border border-gray-300 ...">Ohio</td>
      <td class="border border-gray-300 ...">Columbus</td>
    </tr>
    <tr>
      <td class="border border-gray-300 ...">Michigan</td>
      <td class="border border-gray-300 ...">Detroit</td>
    </tr>
  </tbody>
</table>
```

--------------------------------

### Style based on implicit parent state with in-*

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `in-*` variant to style elements based on the state of any parent element without explicitly adding the `group` class to the parent. This provides a simpler way to apply styles based on parent state changes, though `group` offers more control for fine-grained targeting.

```html
<div tabindex="0" class="group">
  <div class="opacity-50 group-focus:opacity-100"><div tabindex="0">
  <div class="opacity-50 in-focus:opacity-100">
    <!-- ... -->
  </div></div>
</div>
```

--------------------------------

### Control Text Color Opacity with Tailwind CSS

Source: https://tailwindcss.com/docs/color

Append `/` followed by a percentage value to color utilities to adjust the text color opacity of an element.

```html
<p class="text-blue-600/100 dark:text-sky-400/100">The quick brown fox...</p><p class="text-blue-600/75 dark:text-sky-400/75">The quick brown fox...</p><p class="text-blue-600/50 dark:text-sky-400/50">The quick brown fox...</p><p class="text-blue-600/25 dark:text-sky-400/25">The quick brown fox...</p>
```

--------------------------------

### Remove Margins and Padding

Source: https://tailwindcss.com/docs/preflight

Preflight removes default margins and padding from most elements to prevent accidental reliance on user-agent styles.

```css
*,::after,::before,::backdrop,::file-selector-button {
  margin: 0;
  padding: 0;
}
```

--------------------------------

### Set border-y color to yellow shades

Source: https://tailwindcss.com/docs/border-color

Utilize the `border-y-{color}-{shade}` utility to apply various shades of yellow to the top and bottom borders. This covers the full spectrum from 50 to 950.

```html
<div class="border-y-yellow-50">...</div>
```

```html
<div class="border-y-yellow-100">...</div>
```

```html
<div class="border-y-yellow-200">...</div>
```

```html
<div class="border-y-yellow-300">...</div>
```

```html
<div class="border-y-yellow-400">...</div>
```

```html
<div class="border-y-yellow-500">...</div>
```

```html
<div class="border-y-yellow-600">...</div>
```

```html
<div class="border-y-yellow-700">...</div>
```

```html
<div class="border-y-yellow-800">...</div>
```

```html
<div class="border-y-yellow-900">...</div>
```

```html
<div class="border-y-yellow-950">...</div>
```

--------------------------------

### Set Border Color with Olive Palette

Source: https://tailwindcss.com/docs/border-color

These utilities allow you to set border colors using the olive color palette. Each class maps to a `border-inline-color` CSS property with an olive-specific oklch value.

```html
`border-x-olive-50`| `border-inline-color: var(--color-olive-50); /* oklch(98.8% 0.003 106.5) */`
```

```html
`border-x-olive-100`| `border-inline-color: var(--color-olive-100); /* oklch(96.6% 0.005 106.5) */`
```

```html
`border-x-olive-200`| `border-inline-color: var(--color-olive-200); /* oklch(93% 0.007 106.5) */`
```

```html
`border-x-olive-300`| `border-inline-color: var(--color-olive-300); /* oklch(88% 0.011 106.6) */`
```

```html
`border-x-olive-400`| `border-inline-color: var(--color-olive-400); /* oklch(73.7% 0.021 106.9) */`
```

```html
`border-x-olive-500`| `border-inline-color: var(--color-olive-500); /* oklch(58% 0.031 107.3) */`
```

```html
`border-x-olive-600`| `border-inline-color: var(--color-olive-600); /* oklch(46.6% 0.025 107.3) */`
```

```html
`border-x-olive-700`| `border-inline-color: var(--color-olive-700); /* oklch(39.4% 0.023 107.4) */`
```

```html
`border-x-olive-800`| `border-inline-color: var(--color-olive-800); /* oklch(28.6% 0.016 107.4) */`
```

--------------------------------

### Styling Odd and Even Table Rows

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `odd` and `even` variants to apply different background colors to odd and even rows in a table.

```html
<table>  <!-- ... -->  <tbody>    {#each people as person}      <!-- Use different background colors for odd and even rows -->      <tr class="odd:bg-white even:bg-gray-50 dark:odd:bg-gray-900/50 dark:even:bg-gray-950">        <td>{person.name}</td>        <td>{person.title}</td>        <td>{person.email}</td>      </tr>    {/each}  </tbody></table>
```

--------------------------------

### Customize accent color theme variables

Source: https://tailwindcss.com/docs/accent-color

Use `--color-*` theme variables to customize accent color utilities. After defining, the custom utility can be used in your markup.

```css
@theme {  --color-regal-blue: #243a5a; }
```

```html
<input class="accent-regal-blue" type="checkbox" />
```

--------------------------------

### Custom Value Positioning

Source: https://tailwindcss.com/docs/top-right-bottom-left

Apply custom positioning values directly using bracket notation, such as `inset-[3px]`. For CSS variables, use the `inset-(<custom-property>)` syntax.

```html
<div class="inset-[3px] ...">  <!-- ... --></div>
```

```html
<div class="inset-(--my-position) ...">  <!-- ... --></div>
```

--------------------------------

### Use Custom Font Size Utility

Source: https://tailwindcss.com/docs/font-size

Once a custom font size variable is defined, the corresponding utility class can be used directly in your HTML markup.

```html
<div class="text-tiny">  <!-- ... --></div>
```

--------------------------------

### Conditional Styling for Property Support

Source: https://tailwindcss.com/docs/hover-focus-and-other-states

Use the `supports-[property-name]` variant to apply styles when a CSS property is supported, without needing to specify a value. This simplifies checks for property availability.

```html
<div class="bg-black/75 supports-backdrop-filter:bg-black/25 supports-backdrop-filter:backdrop-blur ...">  <!-- ... --></div>
```

--------------------------------

### Start-end logical border radius reset and full

Source: https://tailwindcss.com/docs/border-radius

Remove start-end border-radius or apply infinite radius using logical properties.

```css
border-start-end-radius: 0;
```

```css
border-start-end-radius: calc(infinity * 1px);
```

--------------------------------

### CSS for scrollbar-track-sky-950 utility

Source: https://tailwindcss.com/docs/scrollbar-color

Applies the `scrollbar-track-sky-950` utility to set the scrollbar track color to sky-950.

```css
--tw-scrollbar-track: var(--color-sky-950); /* oklch(29.3% 0.066 243.157) */ scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### CSS for scrollbar-track-sky-700 utility

Source: https://tailwindcss.com/docs/scrollbar-color

Applies the `scrollbar-track-sky-700` utility to set the scrollbar track color to sky-700.

```css
--tw-scrollbar-track: var(--color-sky-700); /* oklch(50% 0.134 242.749) */ scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### Applying Negative Letter Spacing in HTML

Source: https://tailwindcss.com/docs/letter-spacing

Use a negative prefix with a custom tracking utility (e.g., `-tracking-2`) to apply negative letter spacing to an element after defining it in the theme.

```html
<p class="-tracking-2">The quick brown fox ...</p>
```

--------------------------------

### Override entire theme namespace in Tailwind CSS

Source: https://tailwindcss.com/docs/theme

Use the asterisk syntax to reset an entire namespace to initial, then define only custom values. All default utilities in that namespace will be removed.

```CSS
@import "tailwindcss";
@theme {
  --color-*: initial;
  --color-white: #fff;
  --color-purple: #3f3cbb;
  --color-midnight: #121063;
  --color-tahiti: #3ab7bf;
  --color-bermuda: #78dcca;
}
```

--------------------------------

### Bottom Border Color Utilities (Orange)

Source: https://tailwindcss.com/docs/border-color

Apply orange colors to the bottom border using shades from 50 to 400.

```html
border-b-orange-50
```

```css
border-bottom-color: var(--color-orange-50); /* oklch(98% 0.016 73.684) */
```

```html
border-b-orange-100
```

```css
border-bottom-color: var(--color-orange-100); /* oklch(95.4% 0.038 75.164) */
```

```html
border-b-orange-200
```

```css
border-bottom-color: var(--color-orange-200); /* oklch(90.1% 0.076 70.697) */
```

```html
border-b-orange-300
```

```css
border-bottom-color: var(--color-orange-300); /* oklch(83.7% 0.128 66.29) */
```

```html
border-b-orange-400
```

```css
border-bottom-color: var(--color-orange-400); /* oklch(75% 0.183 55.934) */
```

--------------------------------

### CSS for `scrollbar-thumb-lime-400`

Source: https://tailwindcss.com/docs/scrollbar-color

Sets the scrollbar thumb color using the lime-400 shade and the track color using a default variable.

```css
--tw-scrollbar-thumb: var(--color-lime-400); /* oklch(84.1% 0.238 128.85) */ scrollbar-color: var(--tw-scrollbar-thumb) var(--tw-scrollbar-track);
```

--------------------------------

### Cropping background to text

Source: https://tailwindcss.com/docs/background-clip

The `bg-clip-text` utility crops an element's background to match the shape of the text. Requires additional text and color utilities.

```html
<p class="bg-linear-to-r from-pink-500 to-violet-500 bg-clip-text text-5xl font-extrabold text-transparent ...">  Hello world</p>
```

--------------------------------

### Custom object-position value

Source: https://tailwindcss.com/docs/object-position

Use the object-[<value>] syntax to set the object position based on a completely custom value, such as percentages or pixel values. This provides flexibility for precise alignment.

```html
<img class="object-[25%_75%] ..." src="/img/mountains.jpg" />
```