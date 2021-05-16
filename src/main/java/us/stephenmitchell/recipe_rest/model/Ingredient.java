package us.stephenmitchell.recipe_rest.model;

import javax.persistence.*;

@Entity
@Table(name = "ingredient")
public class Ingredient {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "ID")
    private Long id;

    @Column(name="NAME")
    private String name;

    @Column(name="MEASUREMENT")
    private String measurement;

    @Column(name = "GRAMS")
    private Long grams;

    public Ingredient() {
        super();
    }

    public Ingredient(Long id, String name, String measurement, Long grams) {
        super();
        this.id = id;
        this.name = name;
        this.measurement = measurement;
        this.grams = grams;
    }

    public Long getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getMeasurement() {
        return measurement;
    }

    public void setMeasurement(String measurement) {
        this.measurement = measurement;
    }

    public Long getGrams() {
        return grams;
    }

    public void setGrams(Long grams) {
        this.grams = grams;
    }

    @Override
    public String toString() {
        return "Ingredient{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", measurement='" + measurement + '\'' +
                ", grams=" + grams +
                '}';
    }
}
