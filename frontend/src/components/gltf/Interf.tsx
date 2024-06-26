// @ts-nocheck
/*
Auto-generated by: https://github.com/pmndrs/gltfjsx
Command: npx gltfjsx@6.2.18 .\exam.glb -tT 
Files: .\exam.glb [16.12MB] > C:\Users\Christian Kuklis\Code\GitHub\Ki365\Ki365\frontend\src\components\gltf\exam-transformed.glb [676.27KB] (96%)
Pcb_name: exam
Source_pcb_file: exam.glb
Generator: KiCad 8.99.0-unknown
Generated_at: 2024-06-23T07:39:11+0000
*/

import * as THREE from 'three'
import { useGLTF } from '@react-three/drei'
import { GLTF } from 'three-stdlib'

type GLTFResult = GLTF & {
  nodes: {
    C6: THREE.Mesh
    ['=>01113']: THREE.Mesh
    ['=>01116']: THREE.Mesh
    ['=>01117']: THREE.Mesh
    ['=>01118']: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_1: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_2: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_3: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_4: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_5: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_6: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_7: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_8: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_9: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_10: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_11: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_12: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_13: THREE.Mesh
    R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_14: THREE.Mesh
  }
  materials: {
    PaletteMaterial001: THREE.MeshStandardMaterial
    PaletteMaterial002: THREE.MeshStandardMaterial
    PaletteMaterial003: THREE.MeshStandardMaterial
    PaletteMaterial004: THREE.MeshStandardMaterial
  }
}

export function Model(props: JSX.IntrinsicElements['group']) {
  const { nodes, materials } = useGLTF('/interf-transformed.glb') as GLTFResult
  return (
    <group {...props} dispose={null}>
      <mesh geometry={nodes.C6.geometry} material={nodes.C6.material} position={[0.159, 0.002, 0.112]} />
      <mesh geometry={nodes['=>01113'].geometry} material={materials.PaletteMaterial001} />
      <mesh geometry={nodes['=>01116'].geometry} material={materials.PaletteMaterial002} />
      <mesh geometry={nodes['=>01117'].geometry} material={materials.PaletteMaterial003} />
      <mesh geometry={nodes['=>01118'].geometry} material={materials.PaletteMaterial004} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_1.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_1.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_1.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_2.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_2.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_2.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_3.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_3.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_3.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_4.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_4.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_4.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_5.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_5.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_5.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_6.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_6.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_6.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_7.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_7.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_7.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_8.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_8.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_8.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_9.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_9.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_9.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_10.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_10.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_10.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_11.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_11.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_11.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_12.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_12.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_12.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_13.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_13.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_13.instanceMatrix} />
      <instancedMesh args={[nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_14.geometry, nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_14.material, 5]} instanceMatrix={nodes.R_Axial_DIN0207_L63mm_D25mm_P762mm_Horizontal_14.instanceMatrix} />
    </group>
  )
}

useGLTF.preload('/interf-transformed.glb')
