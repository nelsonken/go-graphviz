package graphviz_test

import (
	"bytes"
	"testing"

	"github.com/nelsonken/go-graphviz"
)

func TestGraphviz_Image(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e.SetLabel("e")

	t.Run("png", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
				t.Fatalf("%+v", err)
			}
			if len(buf.Bytes()) != 4602 {
				t.Fatalf("failed to encode png: bytes length is %d", len(buf.Bytes()))
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(graph)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			bounds := image.Bounds()
			if bounds.Max.X != 83 {
				t.Fatal("failed to get image")
			}
			if bounds.Max.Y != 177 {
				t.Fatal("failed to get image")
			}
		})
	})
	t.Run("jpg", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(graph, graphviz.JPG, &buf); err != nil {
				t.Fatalf("%+v", err)
			}
			if len(buf.Bytes()) != 3296 {
				t.Fatalf("failed to encode jpg: bytes length is %d", len(buf.Bytes()))
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(graph)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			bounds := image.Bounds()
			if bounds.Max.X != 83 {
				t.Fatal("failed to get image")
			}
			if bounds.Max.Y != 177 {
				t.Fatal("failed to get image")
			}
		})
	})
}
