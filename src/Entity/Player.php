<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Mapping\ClassMetadata;
use Symfony\Component\Validator\Constraints as Assert;

class Player
{
    /** @var string */
    private $name;

    /** @var int */
    private $vbId;

    public static function loadValidatorMetadata(ClassMetadata $metadata)
    {
        $metadata->addPropertyConstraint('name', new Assert\Length([
            'max' => 128,
        ]));

        $metadata->addPropertyConstraint('vbId', new Assert\NotNull);
        $metadata->addPropertyConstraint('vbId', new Assert\Range([
            'min' => 0,
            'max' => 65535,
        ]));
    }

    /**
     * @param int $vbId
     * @param string $name
     */
    public function __construct($vbId, $name)
    {
        $this->setVbId($vbId);
        $this->setName($name);
    }

    public function getName()
    {
        return $this->name;
    }

    /**
     * @param string $name
     */
    public function setName($name)
    {
        $this->name = (string) $name;
    }

    public function getVbId()
    {
        return $this->vbId;
    }

    /**
     * @param int $vbId
     */
    public function setVbId($vbId)
    {
        $this->vbId = (int) $vbId;
    }
}
