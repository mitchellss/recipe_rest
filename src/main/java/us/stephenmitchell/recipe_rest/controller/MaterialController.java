package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.CollectionModel;
import org.springframework.hateoas.EntityModel;
import org.springframework.web.bind.annotation.*;
import us.stephenmitchell.recipe_rest.assembler.MaterialAssembler;
import us.stephenmitchell.recipe_rest.exception.MaterialNotFoundException;
import us.stephenmitchell.recipe_rest.model.Material;
import us.stephenmitchell.recipe_rest.repository.MaterialRepository;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.linkTo;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.methodOn;

@RestController
@RequestMapping("/api")
public class MaterialController {

    @Autowired
    MaterialRepository materialRepository;

    MaterialAssembler materialAssembler;

    MaterialController(MaterialRepository materialRepository, MaterialAssembler materialAssembler) {
        this.materialRepository = materialRepository;
        this.materialAssembler = materialAssembler;
    }


    @GetMapping("/get_material")
    public CollectionModel<EntityModel<Material>> all() {
        List<EntityModel<Material>> recipes = StreamSupport
                .stream(materialRepository.findAll().spliterator(), false)
                .map(materialAssembler::toModel)
                .collect(Collectors.toList());
        return CollectionModel.of(recipes,
                linkTo(methodOn(MaterialController.class).all()).withSelfRel());
    }

    @PostMapping("/post_material")
    public String postMaterial(@RequestBody Material material) {
        materialRepository.save(material);
        return material.toString();
    }

    @GetMapping("/get_material/{id}")
    public EntityModel<Material> one(@PathVariable Long id) {
        Material material= materialRepository.findById(id)
                .orElseThrow(() -> new MaterialNotFoundException(id));

        return materialAssembler.toModel(material);
    }
}
