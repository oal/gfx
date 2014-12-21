// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfxutil

import (
	"image"
	"os"

	"azul3d.org/gfx.v2-dev"
)

var textures = make(map[string]*gfx.Texture, 32)

// OpenTexture opens the named image file, decodes it, and returns a texture
// with that image as it's source. As usual, you will also need to import a
// image decoder, e.g. for png:
//
//  import _ "image/png"
//
// The returned texture will have a MinFilter == LinearMipmapLinear (trilinear
// filtering) a MagFilter == Linear, and Format == DXT1.
//
// If a error is returned it is an IO or image decoding error and a nil shader
// is returned.
//
// Multiple consecutive calls to OpenTexture with the same exact path will
// yield the same exact texture pointer as a result (they are cached).
func OpenTexture(path string) (*gfx.Texture, error) {
	// If the texture is in the cache already, return that one.
	tex, ok := textures[path]
	if ok {
		return tex, nil
	}

	// Open the image file and decode it.
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// Create the new texture and set it's source to the decoded image.
	tex = gfx.NewTexture()
	tex.Source = img
	tex.Bounds = img.Bounds()

	// Set the texture options.
	tex.MinFilter = gfx.LinearMipmapLinear
	tex.MagFilter = gfx.Linear
	tex.Format = gfx.DXT1

	// Store the texture in the cache for later calls to OpenTexture.
	textures[path] = tex
	return tex, nil
}
