// Author: Ciarán Slemon
// Title: Voronoi

#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

float rand(float index) {
    float BASE_SEED = 12345.776;
    float combinedSeed = BASE_SEED + index;
    return fract(sin(combinedSeed) * 43758.5453123);
}

const int NUM_POINTS = 200; 
vec2 points[NUM_POINTS];

void main() {
    vec2 uv = gl_FragCoord.xy/u_resolution.xy;  
    
    // Create points
    for (int i = 0; i < NUM_POINTS; i++) {
        points[i] = vec2(
            rand(float(i) * 2.0),
        	rand(float(i) * 2.0 + 1.0)            
        );
    }

    // Find closest point
    float shortest = distance(uv.xy, points[0]); 
    float value = 1.0;
        
    for (int i = 0; i < NUM_POINTS; i++) {
        float dist = distance(uv.xy, points[i]);
        if (dist < shortest) {
            shortest = dist;
            value = float(i);
        }
    }
    
    // Create time-based shifting colour
    float mult = abs(sin(u_time + value));
   	vec3 rotatingColour = vec3(0.3 * abs(sin(u_time)), 0.1, 0.8 * mult);
    
    // Create Voronoi cell and reduce brightness
    vec3 base = rotatingColour*(value/30.0);
    
    gl_FragColor = vec4(base, 1.0);    
}

